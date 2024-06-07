package chainlib

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcInterfaceMessages"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcclient"
	"github.com/lavanet/lava/protocol/chainlib/extensionslib"
	"github.com/lavanet/lava/protocol/chaintracker"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/protocol/lavaprotocol"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/protocopy"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type relayFinalizationBlocksHandler interface {
	GetParametersForRelayDataReliability(
		ctx context.Context,
		request *pairingtypes.RelayRequest,
		chainMsg ChainMessage,
		relayTimeout time.Duration,
		blockLagForQosSync int64,
		averageBlockTime time.Duration,
		blockDistanceToFinalization,
		blocksInFinalizationData uint32,
	) (latestBlock int64, requestedBlockHash []byte, requestedHashes []*chaintracker.BlockStore, modifiedReqBlock int64, finalized, updatedChainMessage bool, err error)

	BuildRelayFinalizedBlockHashes(
		ctx context.Context,
		request *pairingtypes.RelayRequest,
		reply *pairingtypes.RelayReply,
		latestBlock int64,
		requestedHashes []*chaintracker.BlockStore,
		updatedChainMessage bool,
		relayTimeout time.Duration,
		averageBlockTime time.Duration,
		blockDistanceToFinalization uint32,
		blocksInFinalizationData uint32,
		modifiedReqBlock int64,
	) (err error)
}

type connectedConsumerContainer struct {
	consumerChannel   *common.SafeChannelSender[*pairingtypes.RelayReply]
	firstSetupRequest *pairingtypes.RelayRequest
}

type activeSubscription struct {
	cancellableContext           context.Context
	cancellableContextCancelFunc context.CancelFunc
	messagesChannel              chan interface{}
	nodeSubscription             *rpcclient.ClientSubscription
	subscriptionID               string
	firstSetupReply              *pairingtypes.RelayReply
	apiCollection                *spectypes.ApiCollection
	connectedConsumers           map[string]*connectedConsumerContainer // key is consumer address
}

type ProviderNodeSubscriptionManager struct {
	chainRouter                    ChainRouter
	chainParser                    ChainParser
	relayFinalizationBlocksHandler relayFinalizationBlocksHandler
	activeSubscriptions            map[string]*activeSubscription // key is request params hash
	privKey                        *btcec.PrivateKey
	lock                           sync.RWMutex
}

func NewProviderNodeSubscriptionManager(chainRouter ChainRouter, chainParser ChainParser, relayFinalizationBlocksHandler relayFinalizationBlocksHandler, privKey *btcec.PrivateKey) *ProviderNodeSubscriptionManager {
	utils.LavaFormatTrace("NewProviderNodeSubscriptionManager")
	return &ProviderNodeSubscriptionManager{
		chainRouter:                    chainRouter,
		chainParser:                    chainParser,
		relayFinalizationBlocksHandler: relayFinalizationBlocksHandler,
		activeSubscriptions:            make(map[string]*activeSubscription),
		privKey:                        privKey,
	}
}

// TODO: Elad: Handle same consumer asking for same subscription twice
func (pnsm *ProviderNodeSubscriptionManager) AddConsumer(ctx context.Context, request *pairingtypes.RelayRequest, chainMessage ChainMessage, consumerAddr sdk.AccAddress, consumerChannel chan<- *pairingtypes.RelayReply) (subscriptionId string, err error) {
	utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() called", utils.LogAttr("consumerAddr", consumerAddr))

	if pnsm == nil {
		return "", fmt.Errorf("ProviderNodeSubscriptionManager is nil")
	}

	hashedParams, params, err := pnsm.getHashedParams(chainMessage)
	if err != nil {
		return "", err
	}

	utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() hashed params",
		utils.LogAttr("params", string(params)),
		utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
	)

	consumerAddrString := consumerAddr.String()

	pnsm.lock.Lock()
	defer pnsm.lock.Unlock()

	var firstSetupReply *pairingtypes.RelayReply

	paramsChannelToConnectedConsumers, found := pnsm.activeSubscriptions[hashedParams]
	if found {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() found existing subscription",
			utils.LogAttr("GUID", ctx),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("params", string(params)),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
		)

		if _, found := paramsChannelToConnectedConsumers.connectedConsumers[consumerAddrString]; found { // Consumer is already connected to this subscription, dismiss
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() consumer is already connected, returning the existing subscription", utils.LogAttr("consumerAddr", consumerAddr))
			return paramsChannelToConnectedConsumers.subscriptionID, nil
		}

		// Add the new entry for the consumer
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() consumer does not exist in the subscription, adding",
			utils.LogAttr("GUID", ctx),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("params", string(params)),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
		)

		paramsChannelToConnectedConsumers.connectedConsumers[consumerAddrString] = &connectedConsumerContainer{
			consumerChannel:   common.NewSafeChannelSender(ctx, consumerChannel),
			firstSetupRequest: &pairingtypes.RelayRequest{}, // Deep copy later	firstSetupChainMessage: chainMessage,
		}

		copyRequestErr := protocopy.DeepCopyProtoObject(request, paramsChannelToConnectedConsumers.connectedConsumers[consumerAddrString].firstSetupRequest)
		if copyRequestErr != nil {
			return "", utils.LavaFormatError("failed to copy subscription request", copyRequestErr)
		}

		firstSetupReply = paramsChannelToConnectedConsumers.firstSetupReply
		subscriptionId = paramsChannelToConnectedConsumers.subscriptionID
	} else {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() did not found existing subscription, creating new one")

		nodeChan := make(chan interface{})
		var replyWrapper *RelayReplyWrapper
		var clientSubscription *rpcclient.ClientSubscription
		// TODO: Elad - change to enum
		replyWrapper, subscriptionId, clientSubscription, _, _, err = pnsm.chainRouter.SendNodeMsg(ctx, nodeChan, chainMessage, append(request.RelayData.Extensions, WebSocketExtension))
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() subscription reply received",
			utils.LogAttr("replyWrapper", replyWrapper),
			utils.LogAttr("subscriptionId", subscriptionId),
			utils.LogAttr("clientSubscription", clientSubscription),
			utils.LogAttr("err", err),
		)
		if err != nil {
			return "", utils.LavaFormatError("ProviderNodeSubscriptionManager: Subscription failed", err, utils.LogAttr("GUID", ctx), utils.LogAttr("params", params))
		}

		if replyWrapper == nil || replyWrapper.RelayReply == nil {
			return "", utils.LavaFormatError("ProviderNodeSubscriptionManager: Subscription failed, relayWrapper or RelayReply are nil", nil, utils.LogAttr("GUID", ctx))
		}

		reply := replyWrapper.RelayReply

		copiedRequest := &pairingtypes.RelayRequest{}
		copyRequestErr := protocopy.DeepCopyProtoObject(request, copiedRequest)
		if copyRequestErr != nil {
			return "", utils.LavaFormatError("failed to copy subscription request", copyRequestErr)
		}

		err = pnsm.signReply(ctx, reply, consumerAddr, chainMessage, request)
		if err != nil {
			return "", err
		}

		if clientSubscription == nil {
			// failed subscription, but not an error. (probably a node error)

			// Send the first message to the consumer, so it can handle the error
			SafeChannelSender := common.NewSafeChannelSender(ctx, consumerChannel)
			SafeChannelSender.Send(reply)

			return "", utils.LavaFormatWarning("ProviderNodeSubscriptionManager: Subscription failed, node error", nil, utils.LogAttr("GUID", ctx), utils.LogAttr("reply", reply))
		}

		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:AddConsumer() subscription successful",
			utils.LogAttr("subscriptionId", subscriptionId),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("firstSetupReplyData", reply.Data),
		)

		cancellableCtx, cancel := context.WithCancel(context.Background())
		channelToConnectedConsumers := &activeSubscription{
			cancellableContext:           cancellableCtx,
			cancellableContextCancelFunc: cancel,
			messagesChannel:              nodeChan,
			nodeSubscription:             clientSubscription,
			subscriptionID:               subscriptionId,
			firstSetupReply:              reply,
			apiCollection:                chainMessage.GetApiCollection(),
			connectedConsumers:           make(map[string]*connectedConsumerContainer),
		}

		channelToConnectedConsumers.connectedConsumers = make(map[string]*connectedConsumerContainer)
		channelToConnectedConsumers.connectedConsumers[consumerAddrString] = &connectedConsumerContainer{
			consumerChannel:   common.NewSafeChannelSender(ctx, consumerChannel),
			firstSetupRequest: copiedRequest,
		}

		pnsm.activeSubscriptions[hashedParams] = channelToConnectedConsumers
		firstSetupReply = reply

		go pnsm.listenForSubscriptionMessages(cancellableCtx, nodeChan, clientSubscription.Err(), hashedParams)
	}

	pnsm.activeSubscriptions[hashedParams].connectedConsumers[consumerAddrString].consumerChannel.Send(firstSetupReply)

	return subscriptionId, nil
}

func (pnsm *ProviderNodeSubscriptionManager) listenForSubscriptionMessages(ctx context.Context, nodeChan chan interface{}, nodeErrChan <-chan error, hashedParams string) {
	utils.LavaFormatTrace("Inside ProviderNodeSubscriptionManager:startListeningForSubscription()", utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)))
	defer utils.LavaFormatTrace("Leaving ProviderNodeSubscriptionManager:startListeningForSubscription()", utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)))

	subscriptionTimeoutTicker := time.NewTicker(15 * time.Minute) // Set a time limit of 15 minutes for the subscription
	defer subscriptionTimeoutTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			// If this context is done, it means that the subscription was already closed
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:startListeningForSubscription() subscription context is done, ending subscription",
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			)
			return
		case <-subscriptionTimeoutTicker.C:
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:startListeningForSubscription() timeout reached, ending subscription",
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			)
			func() {
				pnsm.lock.Lock()
				defer pnsm.lock.Unlock()
				pnsm.closeNodeSubscription(hashedParams)
			}()
			return
		case nodeErr := <-nodeErrChan:
			utils.LavaFormatWarning("ProviderNodeSubscriptionManager:startListeningForSubscription() got error from node, ending subscription", nodeErr,
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			)
			func() {
				pnsm.lock.Lock()
				defer pnsm.lock.Unlock()
				pnsm.closeNodeSubscription(hashedParams)
			}()
			return
		case nodeMsg := <-nodeChan:
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:startListeningForSubscription() got new message from node",
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
				utils.LogAttr("nodeMsg", nodeMsg),
			)
			pnsm.handleNewNodeMessage(ctx, hashedParams, nodeMsg)
		}
	}
}

func (pnsm *ProviderNodeSubscriptionManager) getHashedParams(chainMessage ChainMessageForSend) (hashedParams string, params []byte, err error) {
	rpcInputMessage := chainMessage.GetRPCMessage()
	params, err = json.Marshal(rpcInputMessage.GetParams())
	if err != nil {
		return "", nil, utils.LavaFormatError("could not marshal params", err)
	}

	hashedParams = rpcclient.CreateHashFromParams(params)
	return hashedParams, params, nil
}

func (pnsm *ProviderNodeSubscriptionManager) convertNodeMsgToMarshalledJsonRpcResponse(data interface{}, apiCollection *spectypes.ApiCollection) ([]byte, error) {
	msg, ok := data.(*rpcclient.JsonrpcMessage)
	if !ok {
		return nil, fmt.Errorf("data is not a *rpcclient.JsonrpcMessage, but: %T", data)
	}

	var convertedMsg any
	var err error

	switch apiCollection.GetCollectionData().ApiInterface {
	case spectypes.APIInterfaceJsonRPC:
		convertedMsg, err = rpcInterfaceMessages.ConvertJsonRPCMsg(msg)
		if err != nil {
			return nil, err
		}
	case spectypes.APIInterfaceTendermintRPC:
		convertedMsg, err = rpcInterfaceMessages.ConvertTendermintMsg(msg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported API interface: %s", apiCollection.GetCollectionData().ApiInterface)
	}

	marshalledMsg, err := json.Marshal(convertedMsg)
	if err != nil {
		return nil, err
	}

	return marshalledMsg, nil
}

func (pnsm *ProviderNodeSubscriptionManager) signReply(ctx context.Context, reply *pairingtypes.RelayReply, consumerAddr sdk.AccAddress, chainMessage ChainMessage, request *pairingtypes.RelayRequest) error {
	// Send the first setup message to the consumer in a go routine because the blocking listening for this channel happens after this function
	dataReliabilityEnabled, _ := pnsm.chainParser.DataReliabilityParams()
	blockLagForQosSync, averageBlockTime, blockDistanceToFinalization, blocksInFinalizationData := pnsm.chainParser.ChainBlockStats()
	relayTimeout := GetRelayTimeout(chainMessage, averageBlockTime)

	if dataReliabilityEnabled {
		chainMsgLatestBlock, chainMsgEarliestBlock := chainMessage.RequestedBlock()
		utils.LavaFormatTrace("Before GetParametersForRelayDataReliability",
			utils.LogAttr("chainMsgLatestBlock", chainMsgLatestBlock),
			utils.LogAttr("chainMsgEarliestBlock", chainMsgEarliestBlock),
			utils.LogAttr("replyLatestBlock", reply.LatestBlock),
			utils.LogAttr("requestRequestedBlock", request.RelayData.RequestBlock),
		)
		var err error
		latestBlock, _, requestedHashes, modifiedReqBlock, _, updatedChainMessage, err := pnsm.relayFinalizationBlocksHandler.GetParametersForRelayDataReliability(ctx, request, chainMessage, relayTimeout, blockLagForQosSync, averageBlockTime, blockDistanceToFinalization, blocksInFinalizationData)
		if err != nil {
			return err
		}

		chainMsgLatestBlock, chainMsgEarliestBlock = chainMessage.RequestedBlock()
		utils.LavaFormatTrace("After GetParametersForRelayDataReliability",
			utils.LogAttr("chainMsgLatestBlock", chainMsgLatestBlock),
			utils.LogAttr("chainMsgEarliestBlock", chainMsgEarliestBlock),
			utils.LogAttr("replyLatestBlock", reply.LatestBlock),
			utils.LogAttr("requestRequestedBlock", request.RelayData.RequestBlock),
		)

		err = pnsm.relayFinalizationBlocksHandler.BuildRelayFinalizedBlockHashes(ctx, request, reply, latestBlock, requestedHashes, updatedChainMessage, relayTimeout, averageBlockTime, blockDistanceToFinalization, blocksInFinalizationData, modifiedReqBlock)
		if err != nil {
			return err
		}

		chainMsgLatestBlock, chainMsgEarliestBlock = chainMessage.RequestedBlock()
		utils.LavaFormatTrace("After BuildRelayFinalizedBlockHashes",
			utils.LogAttr("chainMsgLatestBlock", chainMsgLatestBlock),
			utils.LogAttr("chainMsgEarliestBlock", chainMsgEarliestBlock),
			utils.LogAttr("replyLatestBlock", reply.LatestBlock),
			utils.LogAttr("requestRequestedBlock", request.RelayData.RequestBlock),
		)
	}

	var ignoredMetadata []pairingtypes.Metadata
	reply.Metadata, _, ignoredMetadata = pnsm.chainParser.HandleHeaders(reply.Metadata, chainMessage.GetApiCollection(), spectypes.Header_pass_reply)
	reply, err := lavaprotocol.SignRelayResponse(consumerAddr, *request, pnsm.privKey, reply, dataReliabilityEnabled)
	if err != nil {
		return err
	}
	reply.Metadata = append(reply.Metadata, ignoredMetadata...) // appended here only after signing
	return nil
}

func (pnsm *ProviderNodeSubscriptionManager) handleNewNodeMessage(ctx context.Context, hashedParams string, nodeMsg interface{}) {
	pnsm.lock.RLock()
	defer pnsm.lock.RUnlock()

	// Sending message to all connected consumers
	for consumerAddrString, connectedConsumerContainer := range pnsm.activeSubscriptions[hashedParams].connectedConsumers {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:startListeningForSubscription() sending to consumer", utils.LogAttr("consumerAddr", consumerAddrString), utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)))

		copiedRequest := &pairingtypes.RelayRequest{}
		copyRequestErr := protocopy.DeepCopyProtoObject(connectedConsumerContainer.firstSetupRequest, copiedRequest)
		if copyRequestErr != nil {
			utils.LavaFormatError("failed to copy subscription request", copyRequestErr)
			return
		}

		extensionInfo := extensionslib.ExtensionInfo{LatestBlock: 0, ExtensionOverride: copiedRequest.RelayData.Extensions}
		if extensionInfo.ExtensionOverride == nil { // in case consumer did not set an extension, we skip the extension parsing and we are sending it to the regular url
			extensionInfo.ExtensionOverride = []string{}
		}

		chainMessage, err := pnsm.chainParser.ParseMsg(copiedRequest.RelayData.ApiUrl, copiedRequest.RelayData.Data, copiedRequest.RelayData.ConnectionType, copiedRequest.RelayData.GetMetadata(), extensionInfo)
		if err != nil {
			utils.LavaFormatError("failed to parse message", err)
			return
		}

		apiCollection := pnsm.activeSubscriptions[hashedParams].apiCollection

		marshalledNodeMsg, err := pnsm.convertNodeMsgToMarshalledJsonRpcResponse(nodeMsg, apiCollection)
		if err != nil {
			utils.LavaFormatError("error converting node message", err)
			return
		}

		relayMessageFromNode := &pairingtypes.RelayReply{
			Data:     marshalledNodeMsg,
			Metadata: []pairingtypes.Metadata{},
		}

		consumerAddr, err := sdk.AccAddressFromBech32(consumerAddrString)
		if err != nil {
			utils.LavaFormatError("error unmarshalling consumer address", err)
			return
		}

		err = pnsm.signReply(ctx, relayMessageFromNode, consumerAddr, chainMessage, copiedRequest)
		if err != nil {
			utils.LavaFormatError("error signing reply", err)
			return
		}

		utils.LavaFormatDebug("Sending relay to consumer",
			utils.LogAttr("requestRelayData", copiedRequest.RelayData),
			utils.LogAttr("reply", marshalledNodeMsg),
			utils.LogAttr("replyLatestBlock", relayMessageFromNode.LatestBlock),
			utils.LogAttr("consumerAddr", consumerAddr),
		)

		connectedConsumerContainer.consumerChannel.Send(relayMessageFromNode)
	}
}

func (pnsm *ProviderNodeSubscriptionManager) RemoveConsumer(ctx context.Context, chainMessage ChainMessageForSend, consumerAddr sdk.AccAddress, closeConsumerChannel bool) error {
	if pnsm == nil {
		return nil
	}

	hashedParams, params, err := pnsm.getHashedParams(chainMessage)
	if err != nil {
		return err
	}

	utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() requested to remove consumer from subscription",
		utils.LogAttr("GUID", ctx),
		utils.LogAttr("consumerAddr", consumerAddr),
		utils.LogAttr("params", params),
		utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
	)

	consumerAddrString := consumerAddr.String()

	pnsm.lock.Lock()
	defer pnsm.lock.Unlock()

	openSubscriptions, ok := pnsm.activeSubscriptions[hashedParams]
	if !ok {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() no subscription found for params, subscription is already closed",
			utils.LogAttr("GUID", ctx),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("params", params),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
		)
		return nil
	}

	// Remove consumer from connected consumers
	if _, ok := openSubscriptions.connectedConsumers[consumerAddrString]; ok {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() found consumer connected consumers",
			utils.LogAttr("GUID", ctx),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("params", params),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			utils.LogAttr("connectedConsumers", openSubscriptions.connectedConsumers),
		)

		if closeConsumerChannel {
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() closing consumer channel",
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("consumerAddr", consumerAddr),
				utils.LogAttr("params", params),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			)
			openSubscriptions.connectedConsumers[consumerAddrString].consumerChannel.Close()
		}

		delete(pnsm.activeSubscriptions[hashedParams].connectedConsumers, consumerAddrString)
		if len(pnsm.activeSubscriptions[hashedParams].connectedConsumers) == 0 {
			utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() no more connected consumers",
				utils.LogAttr("GUID", ctx),
				utils.LogAttr("consumerAddr", consumerAddr),
				utils.LogAttr("params", params),
				utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			)
			// Cancel the subscription's context and close the subscription
			pnsm.activeSubscriptions[hashedParams].cancellableContextCancelFunc()
			pnsm.closeNodeSubscription(hashedParams)
		}
	} else {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() consumer not found in connected consumers",
			utils.LogAttr("GUID", ctx),
			utils.LogAttr("consumerAddr", consumerAddr),
			utils.LogAttr("params", params),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
			utils.LogAttr("connectedConsumers", openSubscriptions.connectedConsumers),
		)
	}

	utils.LavaFormatTrace("ProviderNodeSubscriptionManager:RemoveConsumer() removed consumer", utils.LogAttr("consumerAddr", consumerAddr), utils.LogAttr("params", params))
	return nil
}

func (pnsm *ProviderNodeSubscriptionManager) closeNodeSubscription(hashedParams string) error {
	if _, ok := pnsm.activeSubscriptions[hashedParams]; !ok {
		return utils.LavaFormatError("closeNodeSubscription called with hashedParams that does not exist", nil, utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)))
	}

	// Disconnect all connected consumers
	for consumerAddrString, consumerChannel := range pnsm.activeSubscriptions[hashedParams].connectedConsumers {
		utils.LavaFormatTrace("ProviderNodeSubscriptionManager:closeNodeSubscription() closing consumer channel",
			utils.LogAttr("consumerAddr", consumerAddrString),
			utils.LogAttr("hashedParams", utils.ToHexString(hashedParams)),
		)
		consumerChannel.consumerChannel.Close()
	}

	pnsm.activeSubscriptions[hashedParams].nodeSubscription.Unsubscribe()
	close(pnsm.activeSubscriptions[hashedParams].messagesChannel)
	delete(pnsm.activeSubscriptions, hashedParams)

	return nil
}

package types

import (
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/sigs"
)

func (rs RelaySession) GetSignature() []byte {
	return rs.Sig
}

func (rs RelaySession) DataToSign() []byte {
	rs.Badge = nil // its not a part of the signature, its a separate part
	rs.Sig = nil
	if len(rs.UnresponsiveProviders) > 0 {
		utils.LavaFormatError("DEBUG unresponsive~!", nil, utils.Attribute{"RelayString", rs.UnresponsiveProviders})
	}
	utils.LavaFormatError("DEBUG", nil, utils.Attribute{"RelayString", rs.String()})
	return []byte(rs.String())
}

func (rs RelaySession) HashRounds() int {
	return 1
}

func (rs RelaySession) CalculateHashForFinalization() []byte {
	sessionIdBytes := sigs.EncodeUint64(rs.SessionId)
	blockHeightBytes := sigs.EncodeUint64(uint64(rs.Epoch))
	relayNumBytes := sigs.EncodeUint64(rs.RelayNum)
	msgParts := [][]byte{
		sessionIdBytes,
		blockHeightBytes,
		relayNumBytes,
	}
	return sigs.Join(msgParts)
}

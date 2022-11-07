package relayer

import (
	context "context"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
	zerologlog "github.com/rs/zerolog/log"
	"log"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/lavanet/lava/relayer/chainproxy"
	"github.com/lavanet/lava/relayer/sentry"
	"github.com/lavanet/lava/relayer/sigs"
	"github.com/lavanet/lava/utils"
	"github.com/spf13/pflag"
)

func PortalServer(
	ctx context.Context,
	clientCtx client.Context,
	listenAddr string,
	chainID string,
	apiInterface string,
	flagSet *pflag.FlagSet,
) {
	//
	rand.Seed(time.Now().UnixNano())
	sk, _, err := utils.GetOrCreateVRFKey(clientCtx)
	if err != nil {
		log.Fatalln("error: GetOrCreateVRFKey", err)
	}
	// Start sentry
	sentry := sentry.NewSentry(clientCtx, chainID, true, nil, nil, apiInterface, sk, flagSet, 0)
	err = sentry.Init(ctx)
	if err != nil {
		log.Fatalln("error sentry.Init", err)
	}
	go sentry.Start(ctx)
	for sentry.GetBlockHeight() == 0 {
		time.Sleep(1 * time.Second)
	}
	g_sentry = sentry
	g_serverChainID = chainID

	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	NEW_RELIC_APP_NAME := myEnv["NEW_RELIC_APP_NAME"]
	NEW_RELIC_LICENSE_KEY := myEnv["NEW_RELIC_LICENSE_KEY"]

	newrelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(NEW_RELIC_APP_NAME),
		newrelic.ConfigLicense(NEW_RELIC_LICENSE_KEY),
		newrelic.ConfigAppLogEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
		func(config *newrelic.Config) {
			zerologlog.Debug().Enabled()
		},
	)
	if myEnv == nil {
		newrelicApp = nil
	}

	//
	// Node
	chainProxy, err := chainproxy.GetChainProxy("", 1, sentry, newrelicApp)
	if err != nil {
		log.Fatalln("error: GetChainProxy", err)
	}

	//
	// Set up a connection to the server.
	log.Printf("PortalServer %s\n", apiInterface)
	keyName, err := sigs.GetKeyName(clientCtx)
	if err != nil {
		log.Fatalln("error: getKeyName", err)
	}
	privKey, err := sigs.GetPrivKey(clientCtx, keyName)
	if err != nil {
		log.Fatalln("error: getPrivKey", err)
	}
	clientKey, _ := clientCtx.Keyring.Key(keyName)
	log.Println("Client pubkey", clientKey.GetPubKey().Address())

	//
	//
	chainProxy.PortalStart(ctx, privKey, listenAddr)
}

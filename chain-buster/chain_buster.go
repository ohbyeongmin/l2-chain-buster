package chainbuster

import (
	"context"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/urfave/cli/v2"
)

// import (
// 	"github.com/ethereum-optimism/optimism/op-service/dial"
// 	"github.com/ethereum-optimism/optimism/op-service/httputil"
// 	oplog "github.com/ethereum-optimism/optimism/op-service/log"
// 	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
// 	"github.com/ethereum-optimism/optimism/op-service/txmgr"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/urfave/cli/v2"
// )

// type Client struct {
// 	TxMgr  txmgr.TxManager
// 	Client *ethclient.Client
// }

// type Buster struct {
// 	Metrics Metricer
// 	Client  *Client

// 	metricsSrv *httputil.HTTPServer
// }

// func NewBuster(cliCtx *cli.Context) (*Buster, error) {
// 	metrics := NewMetrics()
// 	metricsSrv, err := opmetrics.StartServer(metrics.registry, "127.0.0.1", 7344)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logCfg := oplog.ReadCLIConfig(cliCtx)
// 	logger := oplog.NewLogger(oplog.AppOut(cliCtx), logCfg)
// 	txmgrCfg := txmgr.ReadCLIConfig(cliCtx)
// 	txmgrCfg.L1RPCURL = "http://127.0.0.1:9545"

// 	txM, _ := txmgr.NewSimpleTxManager("l2-chanin-buster", logger, metrics, txmgrCfg)
// 	clt, _ := dial.DialEthClientWithTimeout(cliCtx.Context, dial.DefaultDialTimeout, logger, "http://127.0.0.1:9545")

// 	clie := &Client{
// 		TxMgr:  txM,
// 		Client: clt,
// 	}

// 	return &Buster{
// 		Metrics:    metrics,
// 		metricsSrv: metricsSrv,
// 		Client:     clie,
// 	}, nil
// }

func Main() cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, close context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		cfg := NewConfig(cliCtx)
		l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(l.Handler())
		opservice.ValidateEnvVars(flags.EnvPrefix, flags.Flags, l)

		l.Info("Initializing Chain Buster")
		return ChainBusterServiceFromCLIConfig(cliCtx.Context, cfg, l)
	}
}

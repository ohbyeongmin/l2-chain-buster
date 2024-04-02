package chainbuster

import (
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"
	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	L1EthRpc string
	L2EthRpc string
	Scenario string
	Wallets  string

	TxMgrConfig   txmgr.CLIConfig
	LogConfig     oplog.CLIConfig
	MetricsConfig opmetrics.CLIConfig
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		L1EthRpc: ctx.String(flags.L1EthRpcFlag.Name),
		L2EthRpc: ctx.String(flags.L2EthRpcFlag.Name),
		Scenario: ctx.String(flags.ScenarioFlag.Name),
		Wallets:  ctx.String(flags.WalletsFlag.Name),

		TxMgrConfig:   txmgr.ReadCLIConfig(ctx),
		LogConfig:     oplog.ReadCLIConfig(ctx),
		MetricsConfig: opmetrics.ReadCLIConfig(ctx),
	}
}

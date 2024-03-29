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

	TxMgrConfig   txmgr.CLIConfig
	LogConfig     oplog.CLIConfig
	MetricsConfig opmetrics.CLIConfig
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		L1EthRpc: ctx.String(flags.L1EthRpcFlag.Name),
		L2EthRpc: ctx.String(flags.L2EthRpcFlag.Name),

		TxMgrConfig:   txmgr.ReadCLIConfig(ctx),
		LogConfig:     oplog.ReadCLIConfig(ctx),
		MetricsConfig: opmetrics.ReadCLIConfig(ctx),
	}
}

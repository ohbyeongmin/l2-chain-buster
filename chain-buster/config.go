package chainbuster

import (
	"github.com/urfave/cli/v2"

	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"
	"github.com/ohbyeongmin/l2-chain-buster/utils"

	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
)

type YAMLConfig struct {
	Scenarios *Scenarios `yaml:"jobs"`
	Wallets   *Wallets   `yaml:"actors"`
}

type CLIConfig struct {
	L1EthRpc     string
	L2EthRpc     string
	ScenarioPath string

	TxMgrConfig   txmgr.CLIConfig
	LogConfig     oplog.CLIConfig
	MetricsConfig opmetrics.CLIConfig
}

func NewCLIConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		L1EthRpc:     ctx.String(flags.L1EthRpcFlag.Name),
		L2EthRpc:     ctx.String(flags.L2EthRpcFlag.Name),
		ScenarioPath: ctx.String(flags.ScenarioFlag.Name),

		TxMgrConfig:   txmgr.ReadCLIConfig(ctx),
		LogConfig:     oplog.ReadCLIConfig(ctx),
		MetricsConfig: opmetrics.ReadCLIConfig(ctx),
	}
}

func NewYAMLConfig(cfg *CLIConfig) *YAMLConfig {
	var yConfig YAMLConfig
	utils.ConvertYAMLtoStruct(cfg.ScenarioPath, &yConfig)
	return &yConfig
}

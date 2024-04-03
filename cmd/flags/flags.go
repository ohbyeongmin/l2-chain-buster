package flags

import (
	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli/v2"
)

const EnvPrefix = "L2_CHAIN_BUSTER"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvPrefix, name)
}

var (
	L1EthRpcFlag = &cli.StringFlag{
		Name:    "l1-eth-rpc",
		Usage:   "RPC URL for L1",
		EnvVars: prefixEnvVars("L1_ETH_RPC"),
	}
	L2EthRpcFlag = &cli.StringFlag{
		Name:    "l2-eth-rpc",
		Usage:   "RPC URL for L2 geth",
		EnvVars: prefixEnvVars("L2_ETH_RPC"),
	}
	ScenarioFlag = &cli.StringFlag{
		Name:    "scenario-file",
		Usage:   "Scenario file file",
		EnvVars: prefixEnvVars("SCENARIO_PATH"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
	ScenarioFlag,
}

var optionalFlags []cli.Flag

func init() {
	optionalFlags = append(optionalFlags, txmgr.CLIFlags(EnvPrefix)...)
	optionalFlags = append(optionalFlags, opmetrics.CLIFlags(EnvPrefix)...)
	optionalFlags = append(optionalFlags, oplog.CLIFlags(EnvPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag

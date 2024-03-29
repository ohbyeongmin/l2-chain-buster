package flags

import (
	opservice "github.com/ethereum-optimism/optimism/op-service"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli/v2"
)

const envPrefix = "L2_CHAIN_BUSTER"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(envPrefix, name)
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
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
}

var externalFlags []cli.Flag

func init() {
	externalFlags = append(externalFlags, txmgr.CLIFlags(envPrefix)...)
	externalFlags = append(externalFlags, opmetrics.CLIFlags(envPrefix)...)

	Flags = append(requiredFlags, externalFlags...)
}

var Flags []cli.Flag

package chainbuster

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

func Main() cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, close context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		cfg := NewCLIConfig(cliCtx)
		ycfg := NewYAMLConfig(cfg)
		l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(l.Handler())
		opservice.ValidateEnvVars(flags.EnvPrefix, flags.Flags, l)

		l.Info("Initializing Chain Buster")
		return ChainBusterServiceFromConfigs(cliCtx.Context, cfg, ycfg, l)
	}
}

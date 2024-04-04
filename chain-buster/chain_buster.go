package chainbuster

import (
	"context"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/urfave/cli/v2"
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

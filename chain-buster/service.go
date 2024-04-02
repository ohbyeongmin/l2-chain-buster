package chainbuster

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ohbyeongmin/l2-chain-buster/metrics"
)

type ChainBusterService struct {
	Log     log.Logger
	Metrics metrics.Metrics
	Root    *txmgr.TxManager

	ChainBuster

	stopped atomic.Bool
}

func ChainBusterServiceFromCLIConfig(ctx context.Context, cfg *CLIConfig, log log.Logger) (*ChainBusterService, error) {
	var cbs ChainBusterService
	if err := cbs.initFromCLIConfig(ctx, cfg, log); err != nil {
		return nil, errors.Join(err, cbs.Stop(ctx))
	}
	return &cbs, nil
}

func (cbs *ChainBusterService) initFromCLIConfig(ctx context.Context, cfg *CLIConfig, log log.Logger) error {
	// init chainBuster
	if err := cbs.initScenario(cfg); err != nil {
		return fmt.Errorf("failed to init scenario: %w", err)
	}

	fmt.Printf("%+v\n", cbs.Scenarios)

	return nil
}

func (cbs *ChainBusterService) initScenario(cfg *CLIConfig) error {
	new, err := NewScenarios(cfg.Scenario)
	if err != nil {
		return fmt.Errorf("failed to create new scenario: %w", err)
	}
	cbs.Scenarios = new
	return nil
}

func (cbs *ChainBusterService) Start(_ context.Context) error {
	return nil
}

func (cbs *ChainBusterService) Stopped() bool {
	return cbs.stopped.Load()
}

func (cbs *ChainBusterService) Stop(ctx context.Context) error {
	return nil
}
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
	Wallets *Wallets

	ChainBuster

	stopped atomic.Bool
}

func ChainBusterServiceFromConfigs(ctx context.Context, cfg *CLIConfig, ycfg *YAMLConfig, log log.Logger) (*ChainBusterService, error) {
	var cbs ChainBusterService
	if err := cbs.initFromConfigs(ctx, cfg, ycfg, log); err != nil {
		return nil, errors.Join(err, cbs.Stop(ctx))
	}
	fmt.Printf("%+v\n", cbs.Scenarios)
	fmt.Printf("%+v, %d\n", cbs.Wallets, len(cbs.Wallets.List))
	fmt.Printf("%+v\n", ycfg.Root)

	return &cbs, nil
}

func (cbs *ChainBusterService) initFromConfigs(ctx context.Context, cfg *CLIConfig, ycfg *YAMLConfig, log log.Logger) error {
	if err := cbs.initScenario(ycfg); err != nil {
		return fmt.Errorf("failed to init scenario: %w", err)
	}

	if err := cbs.initWallets(cfg, ycfg); err != nil {
		return fmt.Errorf("failed to init wallets: %w", err)
	}

	if err := cbs.initUsers(cfg); err != nil {
		return fmt.Errorf("failed to init users: %w", err)
	}

	return nil
}

func (cbs *ChainBusterService) initScenario(ycfg *YAMLConfig) error {
	if err := CheckScenarios(ycfg.Scenarios); err != nil {
		return fmt.Errorf("faild to init scenario: %w", err)
	}
	cbs.Scenarios = ycfg.Scenarios
	return nil
}

func (cbs *ChainBusterService) initWallets(cfg *CLIConfig, ycfg *YAMLConfig) error {
	if !CheckWallets(ycfg) {
		if err := NewWallets(cfg, ycfg); err != nil {
			return fmt.Errorf("faild to init wallets: %w", err)
		}
	}
	cbs.Wallets = ycfg.Wallets
	return nil
}

func (cbs *ChainBusterService) initRoot(ycfg *YAMLConfig) error {
	return nil
}

func (cbs *ChainBusterService) initUsers(cfg *CLIConfig) error {
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

package chainbuster

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/ohbyeongmin/l2-chain-buster/metrics"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type ChainBusterService struct {
	Log     log.Logger
	Metrics metrics.Metricer
	Root    txmgr.TxManager
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

	return &cbs, nil
}

func (cbs *ChainBusterService) initFromConfigs(ctx context.Context, cfg *CLIConfig, ycfg *YAMLConfig, log log.Logger) error {
	cbs.Log = log
	cbs.initMetrics(cfg)

	if err := cbs.initScenario(ycfg); err != nil {
		return fmt.Errorf("failed to init scenario: %w", err)
	}

	if err := cbs.initWallets(cfg, ycfg); err != nil {
		return fmt.Errorf("failed to init wallets: %w", err)
	}

	if err := cbs.initUsers(cfg); err != nil {
		return fmt.Errorf("failed to init users: %w", err)
	}

	if err := cbs.initRoot(cfg); err != nil {
		return fmt.Errorf("failed to init root: %w", err)
	}

	addr := cbs.Root.From()
	bal, nil := cbs.Client.BalanceAt(ctx, addr, nil)
	fmt.Printf("addr: %v, balance: %d\n", addr, bal)

	return nil
}

func (bs *ChainBusterService) initMetrics(cfg *CLIConfig) {
	if cfg.MetricsConfig.Enabled {
		procName := "default"
		bs.Metrics = metrics.NewMetrics(procName)
	}
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

func (cbs *ChainBusterService) initRoot(cfg *CLIConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := ethclient.DialContext(ctx, cfg.L1EthRpc)
	if err != nil {
		return err
	}
	cbs.Client = client

	root, err := txmgr.NewSimpleTxManager("root", cbs.Log, cbs.Metrics, cfg.TxMgrConfig)
	if err != nil {
		return nil
	}
	cbs.Root = root

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

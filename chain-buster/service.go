package chainbuster

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/ohbyeongmin/l2-chain-buster/metrics"
	"github.com/ohbyeongmin/l2-chain-buster/utils"

	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type Root struct {
	L1TxManager txmgr.TxManager
	L2TxManager txmgr.TxManager
}

type Client struct {
	L1 *ethclient.Client
	L2 *ethclient.Client
}

type ChainBusterService struct {
	Log       log.Logger
	Metrics   metrics.Metricer
	Root      *Root
	Wallets   *Wallets
	Scenarios *Scenarios
	Client    *Client

	driver *ChainBuster

	stopped atomic.Bool
}

func ChainBusterServiceFromConfigs(ctx context.Context, cfg *CLIConfig, ycfg *YAMLConfig, log log.Logger) (*ChainBusterService, error) {
	var cbs ChainBusterService
	if err := cbs.initFromConfigs(ctx, cfg, ycfg, log); err != nil {
		return nil, errors.Join(err, cbs.Stop(ctx))
	}
	fmt.Printf("%+v\n", cbs.Scenarios)
	fmt.Printf("%+v, %d\n", cbs.Wallets, len(cbs.Wallets.List))
	// l1Addr := cbs.Root.L1TxManager.From()
	// l2Addr := cbs.Root.L2TxManager.From()
	// l1bal, _ := cbs.Client.L1.BalanceAt(ctx, l1Addr, nil)
	// l2bal, _ := cbs.Client.L2.BalanceAt(ctx, l2Addr, nil)
	// fmt.Printf("addr: %v, balance: %d\n", l1Addr, l1bal)
	// fmt.Printf("addr: %v, balance: %d\n", l2Addr, l2bal)

	return &cbs, nil
}

func (cbs *ChainBusterService) initFromConfigs(ctx context.Context, cfg *CLIConfig, ycfg *YAMLConfig, log log.Logger) error {
	cbs.Log = log
	cbs.initMetrics(cfg)

	if err := cbs.initScenarios(ycfg); err != nil {
		return fmt.Errorf("failed to init scenarios: %w", err)
	}

	if err := cbs.initWallets(cfg, ycfg); err != nil {
		return fmt.Errorf("failed to init wallets: %w", err)
	}

	if err := cbs.initRoot(cfg); err != nil {
		return fmt.Errorf("failed to init root: %w", err)
	}

	if err := cbs.initClient(ctx, cfg); err != nil {
		return fmt.Errorf("failed to init client: %w", err)
	}

	cbs.initDriver()

	return nil
}

func (bs *ChainBusterService) initMetrics(cfg *CLIConfig) {
	if cfg.MetricsConfig.Enabled {
		procName := "default"
		bs.Metrics = metrics.NewMetrics(procName)
	}
}

func (cbs *ChainBusterService) initScenarios(ycfg *YAMLConfig) error {
	if err := CheckScenarios(ycfg.Scenarios); err != nil {
		return fmt.Errorf("faild to check scenarios: %w", err)
	}
	cbs.Scenarios = ycfg.Scenarios
	return nil
}

func (cbs *ChainBusterService) initWallets(cfg *CLIConfig, ycfg *YAMLConfig) error {
	requiredWallets := cbs.Scenarios.maxUsers() - len(ycfg.Wallets.List)
	wallets := ycfg.Wallets
	for requiredWallets > 0 {
		w, err := NewWallet()
		if err != nil {
			return fmt.Errorf("failed to init wallets: %w", err)
		}
		wallets.List = append(wallets.List, *w)
		requiredWallets -= 1
	}
	if wallets.Reuse {
		utils.ConvertStructsToYAML(cfg.ScenarioPath, ycfg)
	}
	cbs.Wallets = wallets
	return nil
}

func (cbs *ChainBusterService) initRoot(cfg *CLIConfig) error {
	l1TxMgr, err := txmgr.NewSimpleTxManager("l1-root", cbs.Log, cbs.Metrics, cfg.TxMgrConfig)
	if err != nil {
		return fmt.Errorf("failed to init l1TxMgr: %w", err)
	}

	l2TxMgrConfig := cfg.TxMgrConfig
	l2TxMgrConfig.L1RPCURL = cfg.L2EthRpc
	l2TxMgr, err := txmgr.NewSimpleTxManager("l2-root", cbs.Log, cbs.Metrics, l2TxMgrConfig)
	if err != nil {
		return fmt.Errorf("failed to init l2TxMgr: %w", err)
	}

	cbs.Root = &Root{
		L1TxManager: l1TxMgr,
		L2TxManager: l2TxMgr,
	}

	return nil
}

func (cbs *ChainBusterService) initClient(ctx context.Context, cfg *CLIConfig) error {
	l1Client, err := dial.DialEthClientWithTimeout(ctx, dial.DefaultDialTimeout, cbs.Log, cfg.L1EthRpc)
	if err != nil {
		return fmt.Errorf("failed to dial L1 RPC: %w", err)
	}
	l2Client, err := dial.DialEthClientWithTimeout(ctx, dial.DefaultDialTimeout, cbs.Log, cfg.L2EthRpc)
	if err != nil {
		return fmt.Errorf("failed to dial L2 RPC %w", err)
	}

	cbs.Client = &Client{
		L1: l1Client,
		L2: l2Client,
	}

	return nil
}

func (cbs *ChainBusterService) initDriver() {
	cbs.driver = NewChainBuster(DriverSetup{
		Scenarios: cbs.Scenarios,
		Wallets:   cbs.Wallets,
		L1Client:  cbs.Client.L1,
		L2Client:  cbs.Client.L2,
	})
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

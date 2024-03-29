package chainbuster

import (
	"context"
	"sync/atomic"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ohbyeongmin/l2-chain-buster/metrics"
)

type ChainBusterService struct {
	Log       log.Logger
	Metrics   metrics.Metrics
	TxManager []txmgr.TxManager

	stopped atomic.Bool
}

func ChainBusterServiceFromCLIConfig(ctx context.Context, cfg *CLIConfig, log log.Logger) (*ChainBusterService, error) {
	return nil, nil
}

func (cb *ChainBusterService) Start(_ context.Context) error {
	return nil
}

func (cb *ChainBusterService) Stopped() bool {
	return cb.stopped.Load()
}

func (cb *ChainBusterService) Stop(ctx context.Context) error {
	return nil
}

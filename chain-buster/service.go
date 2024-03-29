package chainbuster

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ohbyeongmin/l2-chain-buster/metrics"
)

type ChainBusterService struct {
	Log     log.Logger
	Metrics metrics.Metrics
}

func ChainBusterServiceFromCLIConfig(ctx context.Context, cfg *CLIConfig, log log.Logger) {

}

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	txmetrics "github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
)

type Metricer interface {
	opmetrics.RPCMetricer
	txmetrics.TxMetricer

	RecordTx(tx *txmgr.TxCandidate)
}

type Metrics struct {
	ns       string
	registry *prometheus.Registry
	factory  opmetrics.Factory

	TxCount prometheus.Counter

	txmetrics.TxMetrics
	opmetrics.RPCMetrics
}

func NewMetrics() *Metrics {
	registry := opmetrics.NewRegistry()
	factory := opmetrics.With(registry)
	ns := "l2_chain_buster"

	return &Metrics{
		ns:         ns,
		registry:   registry,
		factory:    factory,
		TxMetrics:  txmetrics.MakeTxMetrics(ns, factory),
		RPCMetrics: opmetrics.MakeRPCMetrics(ns, factory),
		TxCount: factory.NewCounter(prometheus.CounterOpts{
			Namespace: ns,
			Name:      "tx_count",
			Help:      "Number of finished transactions",
		}),
	}
}

func (m *Metrics) RecordTx(tx *txmgr.TxCandidate) {
	m.TxCount.Add(float64(len(tx.TxData)))
	m.TxCount.Inc()
}

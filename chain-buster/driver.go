package chainbuster

import (
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/ethclient"
)

type User struct {
	*txmgr.SimpleTxManager
	txmgr.TxCandidate
}

type ChainBuster struct {
	Scenarios *Scenarios
	Users     []User
	Client    *ethclient.Client
}

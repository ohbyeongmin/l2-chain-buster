package chainbuster

import "github.com/ethereum-optimism/optimism/op-service/txmgr"

type User struct {
	*txmgr.SimpleTxManager
	txmgr.TxCandidate
}

type ChainBuster struct {
	Scenarios *Scenarios
	Users     []User
}

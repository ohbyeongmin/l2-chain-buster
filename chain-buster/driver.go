package chainbuster

import (
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DriverSetup struct {
	Scenarios *Scenarios
	Wallets   *Wallets
	L1Client  *ethclient.Client
	L2Client  *ethclient.Client
}

type User struct {
	*txmgr.SimpleTxManager
	txmgr.TxCandidate
}

type ChainBuster struct {
	DriverSetup
	Users []User
}

func NewChainBuster(setup DriverSetup) *ChainBuster {
	return &ChainBuster{
		DriverSetup: setup,
		Users:       NewUsers(setup.Wallets),
	}
}

func NewUsers(wallets *Wallets) []User {
	return []User{}
}

func (cb *ChainBuster) SendTxOnL2()       {}
func (cb *ChainBuster) DepositL1ToL2()    {}
func (cb *ChainBuster) WithdrawalL2ToL1() {}

func (cb *ChainBuster) Start() {}

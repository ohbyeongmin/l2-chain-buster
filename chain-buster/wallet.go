package chainbuster

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ohbyeongmin/l2-chain-buster/utils"
)

type Wallets struct {
	Reuse bool     `yaml:"reuse_wallets"`
	List  []Wallet `yaml:"wallets"`
}

type Wallet struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"private_key"`
}

// if true enough wallets
func CheckWallets(ycfg *YAMLConfig) bool {
	if max := ycfg.Scenarios.maxUsers(); len(ycfg.Wallets.List) < max {
		return false
	}
	return true
}

func NewWallets(cfg *CLIConfig, ycfg *YAMLConfig) error {
	wallets := ycfg.Wallets
	need := ycfg.Scenarios.maxUsers() - len(wallets.List)
	for need > 0 {
		w, err := newWallet()
		if err != nil {
			return err
		}
		wallets.List = append(wallets.List, *w)
		need -= 1
	}
	if wallets.Reuse {
		utils.ConvertStructsToYAML(cfg.Scenario, ycfg)
	}
	return nil
}

func newWallet() (*Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey := hexutil.Encode(privateKeyBytes[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("faild to create publicKey on %s", privKey)
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &Wallet{
		Address:    address,
		PrivateKey: privKey,
	}, nil
}

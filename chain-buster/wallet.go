package chainbuster

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ohbyeongmin/l2-chain-buster/utils"
)

/*
key: address
value: private key
*/

type Wallets struct {
	Reuse bool     `yaml:"reuse_wallets"`
	List  []Wallet `yaml:"wallets"`
}

type Wallet struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"private_key"`
}

func NewWallets(filename string, scenarios *Scenarios) (*Wallets, error) {
	var wallets Wallets
	if err := utils.ConvertYAMLtoStruct(filename, &wallets); err != nil {
		return nil, err
	}
	if max := scenarios.maxUsers(); len(wallets.List) < max {
		need := max - len(wallets.List)
		for need > 0 {
			w, err := newWallet()
			if err != nil {
				return nil, err
			}
			wallets.List = append(wallets.List, *w)
			need -= 1
		}
		if wallets.Reuse {
			utils.ConvertStructsToYAML(filename, scenarios, &wallets)
		}
	}
	return &wallets, nil
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

package chainbuster

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallets struct {
	Reuse bool     `yaml:"reuse_wallets"`
	List  []Wallet `yaml:"wallets"`
}

type Wallet struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"private_key"`
}

func NewWallet() (*Wallet, error) {
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

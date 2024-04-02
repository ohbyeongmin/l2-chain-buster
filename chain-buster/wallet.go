package chainbuster

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
key: address
value: private key
*/

type Wallets struct {
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

// type Wallets struct {
// 	Keys map[string]string `json:"keys"`
// }

// const DEFAULT_FILENAME = "wallet"

// func GetWallets(fileName string, requiredUsers int) (*Wallets, error) {
// 	wallets := Wallets{
// 		Keys: map[string]string{},
// 	}

// 	if count := wallets.checkUserCount(requiredUsers); count > 0 {
// 		wallets.generatorWallets(count)
// 	}

// 	wallets.WriteJSON(fileName)

// 	return &wallets, nil
// }

// func (w *Wallets) checkUserCount(requiredUsers int) int {
// 	return requiredUsers - len(w.Keys)
// }

// func (w *Wallets) generatorWallets(count int) error {
// 	for count > 0 {
// 		address, privateKey, err := newWallet()
// 		if err != nil {
// 			return err
// 		}
// 		w.Keys[address] = privateKey
// 	}
// 	return nil
// }

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

// func readJSONtoWallets(fileName string) (*Wallets, error) {
// 	return nil, nil
// }

// func (w *Wallets) WriteJSON(fileName string) error {
// 	keysJSON, err := json.MarshalIndent(w.Keys, "", "")
// 	if err != nil {
// 		return err
// 	}
// 	if err := os.WriteFile(fileName, keysJSON, 0766); err != nil {
// 		return err
// 	}
// 	return nil
// }

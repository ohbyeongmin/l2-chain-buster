package optimism

import "github.com/ethereum-optimism/optimism/op-service/txmgr"

type TxManager struct {
	*txmgr.SimpleTxManager
}

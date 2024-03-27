package chainbuster

import (
	"fmt"

	"github.com/ohbyeongmin/l2-chain-buster/pkg/optimism"

	"go.k6.io/k6/js/modules"
)

type Client struct {
	txManagers       []optimism.TxManager
	ComparisonResult string
}

func (c *Client) IsGreater(a, b int) bool {
	if a > b {
		c.ComparisonResult = fmt.Sprintf("%d is greater than %d", a, b)
		return true
	} else {
		c.ComparisonResult = fmt.Sprintf("%d is NOT greater than %d", a, b)
		return false
	}
}

func init() {
	modules.Register("k6/x/l2-chain-buster", new(Client))
}

// func NewClient() {}

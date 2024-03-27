package main

import (
	_ "github.com/ohbyeongmin/l2-chain-buster/chain-buster"

	k6cmd "go.k6.io/k6/cmd"
)

func main() {
	k6cmd.Execute()
}

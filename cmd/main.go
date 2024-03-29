package main

import (
	"os"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	chainbuster "github.com/ohbyeongmin/l2-chain-buster/chain-buster"
	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
)

func main() {
	app := cli.NewApp()
	app.Name = "l2-chain-buster"
	app.Usage = "L2 overload generator"
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Action = cliapp.LifecycleCmd(chainbuster.Main())
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

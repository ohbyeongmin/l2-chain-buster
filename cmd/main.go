package main

import (
	"os"

	"github.com/ohbyeongmin/l2-chain-buster/cmd/flags"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
)

func main() {
	app := cli.NewApp()
	app.Name = "l2-chain-buster"
	app.Usage = "L2 overload generator"
	app.Flags = flags.Flags
	app.Action = func(ctx *cli.Context) error {
		// _, err := chainbuster.NewBuster(ctx)
		// if err != nil {
		// 	return err
		// }
		// return nil
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

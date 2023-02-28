package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"myethereum/cmd/utils"
	"myethereum/internal/flags"
	"os"
)

const (
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

var (
	nodeFlags = flags.Merge([]cli.Flag{
		utils.IdentityFlag,
	})
	rpcFlags = []cli.Flag{
		utils.HTTPEnabledFlag,
	}
	metricsFlags = []cli.Flag{
		utils.MetricsEnabledFlag,
	}
)
var app = flags.NewApp("the go-ethereum command line interface")

func init() {
	app.Action = geth
	app.HideVersion = true
	app.Copyright = "Copyright The kimin zhang"
	app.Commands = []*cli.Command{
		initCommand,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func geth(ctx *cli.Context) error {
	if args := ctx.Args().Slice(); len(args) > 0 {
		return fmt.Errorf("invalid Command:%q", args[0])
	}
	if ctx.IsSet(utils.Name) {
		fmt.Println("set identity")
	}
}

package utils

import (
	"github.com/urfave/cli/v2"
	"myethereum/internal/flags"
)

var (
	IdentityFlag = &cli.StringFlag{
		Name:     "identity",
		Usage:    "Custom node name",
		Category: flags.NetworkingCategory,
	}
	HTTPEnabledFlag = &cli.BoolFlag{
		Name:     "http",
		Usage:    "Enable the http server",
		Category: flags.APICategory,
	}
	MetricsEnabledFlag = &cli.BoolFlag{
		Name:     "metrics",
		Usage:    "Enable metrics collection and reporting",
		Category: flags.MetricsCategory,
	}
)
var app = flags.NewApp("the go ethereum command line interface")

func init() {
	app.Action = geth
}

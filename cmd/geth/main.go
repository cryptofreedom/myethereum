package geth

import (
	"github.com/urfave/cli/v2"
	"myethereum/cmd/utils"
	"myethereum/internal/flags"
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
		utils.M
	}
)

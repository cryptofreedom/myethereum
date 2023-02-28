package utils

import (
	"github.com/ethereum/go-ethereum/eth/ethconfig"
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
	CachePreimagesFlag = &cli.BoolFlag{
		Name:     "cache.preimages",
		Usage:    "Enable recording the SHA3/keccak preimages of trie keys",
		Category: flags.PerfCategory,
	}
	RemoteDBFlag = &cli.StringFlag{
		Name:     "remotedb",
		Usage:    "URL for remote database",
		Category: flags.LoggingCategory,
	}
	DBEngineFlag = &cli.StringFlag{
		Name:     "db.engine",
		Usage:    "Backing database implementation to use ('leveldb' or 'pebble')",
		Value:    "leveldb",
		Category: flags.EthCategory,
	}
	AncientFlag = &flags.DirectoryFlag{
		Name:     "datadir.ancient",
		Usage:    "Root directory for ancient data (default = inside chaindata)",
		Category: flags.EthCategory,
	}
	NetworkIdFlag = &cli.Uint64Flag{
		Name:     "networkid",
		Usage:    "Explicitly set network id (integer)(For testnets: use --rinkeby, --goerli, --sepolia instead)",
		Value:    ethconfig.Defaults.NetworkId,
		Category: flags.EthCategory,
	}
	// DatabasePathFlags is the flag group of all database path flags.
	DatabasePathFlags = []cli.Flag{
		AncientFlag,
		RemoteDBFlag,
	}
)
var app = flags.NewApp("the go ethereum command line interface")

func init() {
	//rawdb init
}

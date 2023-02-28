package main

import (
	"github.com/urfave/cli/v2"
	"myethereum/cmd/utils"
	"myethereum/internal/flags"
)

var (
	initCommand = &cli.Command{
		Action:    initGenesis,
		Name:      "init",
		Usage:     "Bootstrap and initialize a new genesis block",
		ArgsUsage: "<genesisPath>",
		Flags:     flags.Merge([]cli.Flag{utils.CachePreimagesFlag}, utils.DatabasePathFlags),
		Description: `
The init command initializes a new genesis block and definition for the network.
This is a destructive action and changes the network in which you will be
participating.

It expects the genesis file as argument.`,
	}
)

func initGenesis(ctx *cli.Context) error {
	return nil
}

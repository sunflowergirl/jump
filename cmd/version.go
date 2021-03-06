package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const version = "0.5.1"

func versionCmd(cli.Args, *config.Config) {
	cli.Errf("%s\n", version)
}

func init() {
	cli.RegisterCommand("--version", "Show version.", versionCmd)
}

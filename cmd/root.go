package cmd

import (
	"examplecli/cmd/algorithms"
	"examplecli/cmd/games"
	"examplecli/cmd/tools"

	"github.com/urfave/cli/v2"
)

var (
	RootCmd = &cli.App{
		Commands: []*cli.Command{
			algorithms.RootCmd,
			games.RootCmd,
			tools.RootCmd,
		},
	}
)

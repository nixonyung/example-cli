package tools

import (
	baseconverter "examplecli/cmd/tools/base_converter"
	persongenerator "examplecli/cmd/tools/person_generator"
	queryparam "examplecli/cmd/tools/query_param"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "tools",
	Subcommands: []*cli.Command{
		baseconverter.RootCmd,
		persongenerator.RootCmd,
		queryparam.RootCmd,
	},
}

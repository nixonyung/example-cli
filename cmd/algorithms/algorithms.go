package algorithms

import (
	"examplecli/cmd/algorithms/exponential"
	"examplecli/cmd/algorithms/gcd"
	"examplecli/cmd/algorithms/lcm"
	"examplecli/cmd/algorithms/primes"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "algorithms",
	Subcommands: []*cli.Command{
		exponential.RootCmd,
		gcd.RootCmd,
		lcm.RootCmd,
		primes.RootCmd,
	},
}

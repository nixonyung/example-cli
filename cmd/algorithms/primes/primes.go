package primes

import (
	"errors"
	"examplecli/pkg/algorithms"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func parseArgs(ctx *cli.Context) (int, error) {
	var returnErr error = nil

	// check args length
	if ctx.Args().Len() != 2 {
		// early exit
		return 0, errors.New("requires 2 arguments BASE POWER")
	}

	// check in0
	var in0 int
	if val, err := strconv.Atoi(ctx.Args().Get(0)); err != nil {
		returnErr = errors.Join(returnErr, err)
	} else if val < 0 {
		returnErr = errors.Join(returnErr, errors.New("args[0]: BASE should be a positive integer"))
	} else {
		in0 = val
	}

	return in0, returnErr
}

var RootCmd = &cli.Command{
	Name:      "primes",
	ArgsUsage: "[UPPER_BOUND]",
	Action: func(ctx *cli.Context) error {
		in0, err := parseArgs(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Primes <= %d: %v\n", in0, algorithms.Primes(in0))
		return nil
	},
}

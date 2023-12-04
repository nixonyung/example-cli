package lcm

import (
	"errors"
	"examplecli/pkg/algorithms"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func parseArgs(ctx *cli.Context) (int, int, error) {
	var returnErr error = nil

	// check args length
	if ctx.Args().Len() != 2 {
		// early exit
		return 0, 0, errors.New("requires 2 arguments NUM1 NUM2")
	}

	// check in0
	var in0 int
	if val, err := strconv.Atoi(ctx.Args().Get(0)); err != nil {
		returnErr = errors.Join(returnErr, err)
	} else if val < 0 {
		returnErr = errors.Join(returnErr, errors.New("args[0]: NUM1 should be a positive integer"))
	} else {
		in0 = val
	}

	// check in1
	var in1 int
	if val, err := strconv.Atoi(ctx.Args().Get(1)); err != nil {
		returnErr = errors.Join(returnErr, err)
	} else if val < 0 {
		returnErr = errors.Join(returnErr, errors.New("args[1]: NUM2 should be a positive integer"))
	} else {
		in1 = val
	}

	return in0, in1, returnErr
}

var RootCmd = &cli.Command{
	Name:      "lcm",
	ArgsUsage: "[NUM1] [NUM2]",
	Action: func(ctx *cli.Context) error {
		in0, in1, err := parseArgs(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("LCM of %d and %d is %d\n", in0, in1, algorithms.LCM(in0, in1))
		return nil
	},
}

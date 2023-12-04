package queryparam

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	PATTERN *regexp.Regexp
)

func init() {
	// init PATTERN
	if val, err := regexp.Compile("^\\w+=(\"[[:print:]]*\"|\\d+)(&\\w+=(\"[[:print:]]*\"|\\d+))*$"); err != nil {
		panic(err)
	} else {
		PATTERN = val
	}
}

func parseArgs(ctx *cli.Context) (string, error) {
	var returnErr error = nil

	// check args length
	if ctx.Args().Len() != 1 {
		// early exit
		return "", errors.New("requires 1 arguments INPUT")
	}

	// check in0
	var in0 string = ctx.Args().Get(0)
	if !PATTERN.Match([]byte(in0)) {
		returnErr = errors.Join(returnErr, errors.New("args[0]: wrong format"))
	}

	return in0, returnErr
}

var RootCmd = &cli.Command{
	Name: "query_param",
	Action: func(ctx *cli.Context) error {
		in0, err := parseArgs(ctx)
		if err != nil {
			return err
		}

		result := map[string]any{}
		for _, param := range strings.Split(in0, "&") {
			tokens := strings.SplitN(param, "=", 2)
			key := tokens[0]
			value := tokens[1]
			if strings.HasPrefix(value, "\"") {
				// value is a string (wrapped with double-quotes)
				result[key] = value[1 : len(value)-1]
			} else {
				// else value is an integer
				result[key], _ = strconv.Atoi(value)
			}
		}

		if output, err := json.MarshalIndent(result, "", "  "); err != nil {
			return fmt.Errorf("error when json.MarshalIndent result: %w", err)
		} else {
			fmt.Println(string(output))
		}

		return nil
	},
}

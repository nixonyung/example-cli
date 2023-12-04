package guessing

import (
	"examplecli/pkg/console"
	"examplecli/pkg/env"
	"fmt"
	"math/rand"

	"github.com/urfave/cli/v2"
)

var (
	ANS_MIN = env.GetIntEnv("ANS_MIN", 1)   // inclusive
	ANS_MAX = env.GetIntEnv("ANS_MAX", 100) // inclusive
	ANS     = env.GetIntEnv("ANS", rand.Intn(ANS_MAX-ANS_MIN+1)+ANS_MIN)

	inputMin = ANS_MIN // inclusive
	inputMax = ANS_MAX // inclusive
)

func validateConfig() error {
	if ANS < ANS_MIN || ANS > ANS_MAX {
		return fmt.Errorf(
			"invalid config: ANS is not within ANS_MIN and ANS_MAX, want %d to %d",
			ANS_MIN,
			ANS_MAX,
		)
	}
	return nil
}

var (
	RootCmd = &cli.Command{
		Name: "guessing",
		Action: func(_ *cli.Context) error {
			if err := validateConfig(); err != nil {
				return err
			}

			for {
				fmt.Printf("Enter an integer between %d to %d: ", inputMin, inputMax)
				val, err := console.ScanInt(10)
				if err != nil {
					fmt.Printf("invalid input: %s\n", err.Error())
					continue
				} else if val < inputMin || val > inputMax {
					fmt.Println("invalid range")
					continue
				}

				if val == ANS {
					fmt.Println("bingo!!!")
					break
				} else if val < ANS {
					fmt.Println("too small!")
					inputMin = val + 1
				} else {
					fmt.Println("too large!")
					inputMax = val - 1
				}
			}
			return nil
		},
	}
)

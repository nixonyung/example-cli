package main

import (
	"fmt"
	"os"

	"examplecli/cmd"
)

func main() {
	if err := cmd.RootCmd.Run(os.Args); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

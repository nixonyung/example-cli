package console

import (
	"examplecli/pkg/env"
	"fmt"
	"strconv"
)

var (
	ECHO = env.GetBoolEnv("ECHO", false)
)

func ScanString() (string, error) {
	var line string
	if _, err := fmt.Scanln(&line); err != nil {
		return "", err
	}
	if ECHO {
		fmt.Println(line)
	}
	return line, nil
}

func ScanInt(base int) (int, error) {
	line, err := ScanString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(line, base, 0)
	if err != nil {
		return 0, err
	}

	return int(val), nil
}

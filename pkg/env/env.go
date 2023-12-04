package env

import (
	"os"
	"strconv"
)

func GetBoolEnv(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); !ok {
		return defaultVal
	} else if val == "0" {
		return false
	} else if val == "1" {
		return true
	} else {
		return defaultVal
	}
}

func GetIntEnv(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); !ok {
		return defaultVal
	} else if parsed, err := strconv.Atoi(val); err != nil {
		return defaultVal
	} else {
		return parsed
	}
}

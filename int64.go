package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt64(k string, d int64) int64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize64)
	if err != nil {
		return d
	}

	return i
}

func MustGetInt64(k string) int64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return i
}

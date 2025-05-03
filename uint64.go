package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetUint64(k string, d uint64) uint64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize64)
	if err != nil {
		return d
	}

	return i
}

func MustGetUint64(k string) uint64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return i
}

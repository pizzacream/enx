package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetUint(k string, d uint) uint {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseUint(s, defaultBase, defaultBitSize)
	if err != nil {
		return d
	}

	return uint(i)
}

func MustGetUint(k string) uint {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseUint(s, defaultBase, defaultBitSize)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return uint(i)
}

package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetUint8(k string, d uint8) uint8 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize8)
	if err != nil {
		return d
	}

	return uint8(i)
}

func MustGetUint8(k string) uint8 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return uint8(i)
}

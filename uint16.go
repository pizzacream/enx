package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetUint16(k string, d uint16) uint16 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize16)
	if err != nil {
		return d
	}

	return uint16(i)
}

func MustGetUint16(k string) uint16 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return uint16(i)
}

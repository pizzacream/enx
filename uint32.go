package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetUint32(k string, d uint32) uint32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize32)
	if err != nil {
		return d
	}

	return uint32(i)
}

func MustGetUint32(k string) uint32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseUint(s, defaultBase, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return uint32(i)
}

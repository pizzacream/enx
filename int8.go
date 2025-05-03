package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt8(k string, d int8) int8 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize8)
	if err != nil {
		return d
	}

	return int8(i)
}

func MustGetInt8(k string) int8 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return int8(i)
}

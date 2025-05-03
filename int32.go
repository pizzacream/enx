package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt32(k string, d int32) int32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize32)
	if err != nil {
		return d
	}

	return int32(i)
}

func MustGetInt32(k string) int32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return int32(i)
}

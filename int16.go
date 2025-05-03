package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt16(k string, d int16) int16 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize16)
	if err != nil {
		return d
	}

	return int16(i)
}

func MustGetInt16(k string) int16 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	i, err := strconv.ParseInt(s, defaultBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return int16(i)
}

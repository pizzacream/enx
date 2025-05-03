package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetBool(k string, d bool) bool {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return d
	}

	return b
}

func MustGetBool(k string) bool {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return b
}

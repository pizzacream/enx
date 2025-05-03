package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt(k string, d int) int {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}

	return i
}

func MustGetInt(k string) int {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return v
}

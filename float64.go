package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloat64(k string, d float64) float64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	f, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		return d
	}

	return f
}

func MustGetFloat64(k string) float64 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	f, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return f
}

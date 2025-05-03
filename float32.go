package enx

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloat32(k string, d float32) float32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	f, err := strconv.ParseFloat(s, bitSize32)
	if err != nil {
		return d
	}

	return float32(f)
}

func MustGetFloat32(k string) float32 {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	f, err := strconv.ParseFloat(s, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, s))
	}

	return float32(f)
}

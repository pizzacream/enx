package enx

import (
	"fmt"
	"os"
)

func GetString(k string, d string) string {
	s, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	return s
}

func MustGetString(k string) string {
	s, ok := os.LookupEnv(k)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	return s
}

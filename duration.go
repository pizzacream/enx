package enx

import (
	"fmt"
	"os"
	"time"
)

func GetDuration(k string, d time.Duration) time.Duration {
	value := os.Getenv(k)
	if value == "" {
		return d
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return d
	}

	return duration
}

func MustGetDuration(k string) time.Duration {
	value := os.Getenv(k)
	if value == "" {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, value))
	}

	return duration
}

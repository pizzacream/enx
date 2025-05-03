package enx

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetJSON[T any](key string, fallback T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	var v T
	err := json.Unmarshal([]byte(val), &v)
	if err != nil {
		return fallback
	}

	return v
}

func MustGetJSON[T any](key string) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	var v T
	err := json.Unmarshal([]byte(val), &v)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, val))
	}

	return v
}

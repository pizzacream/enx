package enx

import (
	"fmt"
	"os"
	"regexp"
)

func GetRegex(k string, d *regexp.Regexp) *regexp.Regexp {
	value := os.Getenv(k)
	fmt.Println(value)
	if value == "" {
		return d
	}

	regex, err := regexp.Compile(value)
	if err != nil {
		return d
	}

	return regex
}

func MustGetRegex(k string) *regexp.Regexp {
	value := os.Getenv(k)
	if value == "" {
		panic(fmt.Sprintf("environment variable '%s' not set", k))
	}

	regex, err := regexp.Compile(value)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", k, value))
	}

	return regex
}

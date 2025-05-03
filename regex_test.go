package enx

import (
	"errors"
	"os"
	"regexp"
	"testing"
)

func TestGetRegex(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := regexp.MustCompile(".*")

		os.Setenv("REGEX", ".*")
		defer os.Unsetenv("REGEX")

		got := GetRegex("REGEX", nil)
		if got.String() != expected.String() {
			t.Errorf("GetRegex should return %v for 'REGEX'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		got := GetRegex("REGEX", nil)
		if got != nil {
			t.Errorf("GetRegex should return nil for 'REGEX'")
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		os.Setenv("REGEX", "(")
		defer os.Unsetenv("REGEX")

		got := GetRegex("REGEX", nil)
		if got != nil {
			t.Errorf("GetRegex should return nil for 'REGEX'")
		}
	})
}

func TestMustGetRegex(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := regexp.MustCompile(".*")

		os.Setenv("REGEX", ".*")
		defer os.Unsetenv("REGEX")

		got := MustGetRegex("REGEX")
		if got.String() != expected.String() {
			t.Errorf("MustGetRegex should return %v for 'REGEX'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetRegex: environment variable 'NON_EXISTING_REGEX' not set"))

		MustGetRegex("NON_EXISTING_REGEX")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetRegex: invalid environment variable 'REGEX' has been set: ("))

		os.Setenv("REGEX", "(")
		defer os.Unsetenv("REGEX")

		MustGetRegex("REGEX")
	})
}

package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetBool(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := true

		os.Setenv("TEST_BOOL", "true")
		defer os.Unsetenv("TEST_BOOL")

		got := GetBool("TEST_BOOL", false)

		if got != expected {
			t.Errorf("GetBool should return %v for 'TEST_BOOL'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := true

		got := GetBool("NON_EXISTING_BOOL", true)

		if got != expected {
			t.Errorf("GetBool should return %v for 'NON_EXISTING_BOOL'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := true

		os.Setenv("INVALID_BOOL", "invalid")
		defer os.Unsetenv("INVALID_BOOL")

		got := GetBool("INVALID_BOOL", true)

		if got != expected {
			t.Errorf("GetBool should return %v for 'INVALID_BOOL'", expected)
		}
	})
}

func TestMustGetBool(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := true

		os.Setenv("TEST_BOOL", "true")
		defer os.Unsetenv("TEST_BOOL")

		got := MustGetBool("TEST_BOOL")

		if got != expected {
			t.Errorf("MustGetBool should return %v for 'TEST_BOOL'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetBool: environment variable 'NON_EXISTING_BOOL' not set"))

		MustGetBool("NON_EXISTING_BOOL")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetBool: invalid environment variable 'TEST_BOOL' has been set: invalid"))

		os.Setenv("INVALID_BOOL", "invalid")
		defer os.Unsetenv("INVALID_BOOL")

		MustGetBool("INVALID_BOOL")
	})
}

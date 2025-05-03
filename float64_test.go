package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetFloat64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := 123.456

		os.Setenv("TEST_FLOAT64", "123.456")
		defer os.Unsetenv("TEST_FLOAT64")

		got := GetFloat64("TEST_FLOAT64", 0)

		if got != expected {
			t.Errorf("GetFloat64 should return %v for 'TEST_FLOAT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := 123.456

		got := GetFloat64("NON_EXISTING_FLOAT64", expected)

		if got != expected {
			t.Errorf("GetFloat64 should return %v for 'NON_EXISTING_FLOAT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := 123.456

		os.Setenv("INVALID_FLOAT64", "invalid")
		defer os.Unsetenv("INVALID_FLOAT64")

		got := GetFloat64("INVALID_FLOAT64", expected)

		if got != expected {
			t.Errorf("GetFloat64 should return %v for 'INVALID_FLOAT64'", expected)
		}
	})
}

func TestMustGetFloat64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := 123.456

		os.Setenv("TEST_FLOAT64", "123.456")
		defer os.Unsetenv("TEST_FLOAT64")

		got := MustGetFloat64("TEST_FLOAT64")

		if got != expected {
			t.Errorf("MustGetFloat64 should return %v for 'TEST_FLOAT64'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetFloat64: environment variable 'NON_EXISTING_FLOAT64' not set"))

		MustGetFloat64("NON_EXISTING_FLOAT64")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetFloat64: invalid environment variable 'TEST_FLOAT64' has been set: invalid"))

		os.Setenv("INVALID_FLOAT64", "invalid")
		defer os.Unsetenv("INVALID_FLOAT64")

		MustGetFloat64("INVALID_FLOAT64")
	})
}

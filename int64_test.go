package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetInt64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int64(123)

		os.Setenv("TEST_INT64", "123")
		defer os.Unsetenv("TEST_INT64")

		got := GetInt64("TEST_INT64", 0)

		if got != expected {
			t.Errorf("GetInt64 should return %v for 'TEST_INT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := int64(123)

		got := GetInt64("NON_EXISTING_INT64", expected)

		if got != expected {
			t.Errorf("GetInt64 should return %v for 'NON_EXISTING_INT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := int64(123)

		os.Setenv("INVALID_INT64", "invalid")
		defer os.Unsetenv("INVALID_INT64")

		got := GetInt64("INVALID_INT64", expected)

		if got != expected {
			t.Errorf("GetInt64 should return %v for 'INVALID_INT64'", expected)
		}
	})
}

func TestMustGetInt64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int64(123)

		os.Setenv("TEST_INT64", "123")
		defer os.Unsetenv("TEST_INT64")

		got := MustGetInt64("TEST_INT64")

		if got != expected {
			t.Errorf("MustGetInt64 should return %v for 'TEST_INT64'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt64: environment variable 'NON_EXISTING_INT64' not set"))

		MustGetInt64("NON_EXISTING_INT64")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt64: invalid environment variable 'TEST_INT64' has been set: invalid"))

		os.Setenv("INVALID_INT64", "invalid")
		defer os.Unsetenv("INVALID_INT64")

		MustGetInt64("INVALID_INT64")
	})
}

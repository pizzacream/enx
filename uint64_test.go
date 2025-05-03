package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetUint64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint64(123)

		os.Setenv("TEST_UINT64", "123")
		defer os.Unsetenv("TEST_UINT64")

		got := GetUint64("TEST_UINT64", 0)

		if got != expected {
			t.Errorf("GetUint64 should return %v for 'TEST_UINT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := uint64(123)

		got := GetUint64("NON_EXISTING_UINT64", expected)

		if got != expected {
			t.Errorf("GetUint64 should return %v for 'NON_EXISTING_UINT64'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := uint64(123)

		os.Setenv("INVALID_UINT64", "invalid")
		defer os.Unsetenv("INVALID_UINT64")

		got := GetUint64("INVALID_UINT64", expected)

		if got != expected {
			t.Errorf("GetUint64 should return %v for 'INVALID_UINT64'", expected)
		}
	})
}

func TestMustGetUint64(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint64(123)

		os.Setenv("TEST_UINT64", "123")
		defer os.Unsetenv("TEST_UINT64")

		got := MustGetUint64("TEST_UINT64")

		if got != expected {
			t.Errorf("MustGetUint64 should return %v for 'TEST_UINT64'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint64: environment variable 'NON_EXISTING_UINT64' not set"))

		MustGetUint64("NON_EXISTING_UINT64")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint64: invalid environment variable 'TEST_UINT64' has been set: invalid"))

		os.Setenv("INVALID_UINT64", "invalid")
		defer os.Unsetenv("INVALID_UINT64")

		MustGetUint64("INVALID_UINT64")
	})
}

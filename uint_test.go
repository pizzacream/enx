package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetUint(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint(123)

		os.Setenv("TEST_UINT", "123")
		defer os.Unsetenv("TEST_UINT")

		got := GetUint("TEST_UINT", 0)

		if got != expected {
			t.Errorf("GetUint should return %v for 'TEST_UINT'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := uint(123)

		got := GetUint("NON_EXISTING_UINT", 123)

		if got != expected {
			t.Errorf("GetUint should return %v for 'NON_EXISTING_UINT'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := uint(123)

		os.Setenv("INVALID_UINT", "invalid")
		defer os.Unsetenv("INVALID_UINT")

		got := GetUint("INVALID_UINT", 123)

		if got != expected {
			t.Errorf("GetUint should return %v for 'INVALID_UINT'", expected)
		}
	})
}

func TestMustGetUint(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint(123)

		os.Setenv("TEST_UINT", "123")
		defer os.Unsetenv("TEST_UINT")

		got := MustGetUint("TEST_UINT")

		if got != expected {
			t.Errorf("MustGetUint should return %v for 'TEST_UINT'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint: environment variable 'NON_EXISTING_UINT' not set"))

		MustGetUint("NON_EXISTING_UINT")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint: invalid environment variable 'TEST_UINT' has been set: invalid"))

		os.Setenv("INVALID_UINT", "invalid")
		defer os.Unsetenv("INVALID_UINT")

		MustGetUint("INVALID_UINT")
	})
}

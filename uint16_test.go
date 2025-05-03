package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetUint16(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint16(123)

		os.Setenv("TEST_UINT16", "123")
		defer os.Unsetenv("TEST_UINT16")

		got := GetUint16("TEST_UINT16", 0)

		if got != expected {
			t.Errorf("GetUint16 should return %v for 'TEST_UINT16'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := uint16(123)

		got := GetUint16("NON_EXISTING_UINT16", expected)

		if got != expected {
			t.Errorf("GetUint16 should return %v for 'NON_EXISTING_UINT16'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := uint16(123)

		os.Setenv("INVALID_UINT16", "invalid")
		defer os.Unsetenv("INVALID_UINT16")

		got := GetUint16("INVALID_UINT16", expected)

		if got != expected {
			t.Errorf("GetUint16 should return %v for 'INVALID_UINT16'", expected)
		}
	})
}

func TestMustGetUint16(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint16(123)

		os.Setenv("TEST_UINT16", "123")
		defer os.Unsetenv("TEST_UINT16")

		got := MustGetUint16("TEST_UINT16")

		if got != expected {
			t.Errorf("MustGetUint16 should return %v for 'TEST_UINT16'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint16: environment variable 'NON_EXISTING_UINT16' not set"))

		MustGetUint16("NON_EXISTING_UINT16")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint16: invalid environment variable 'TEST_UINT16' has been set: invalid"))

		os.Setenv("INVALID_UINT16", "invalid")
		defer os.Unsetenv("INVALID_UINT16")

		MustGetUint16("INVALID_UINT16")
	})
}

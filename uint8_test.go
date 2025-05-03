package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetUint8(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint8(123)

		os.Setenv("TEST_UINT8", "123")
		defer os.Unsetenv("TEST_UINT8")

		got := GetUint8("TEST_UINT8", 0)

		if got != expected {
			t.Errorf("GetUint8 should return %v for 'TEST_UINT8'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := uint8(123)

		got := GetUint8("NON_EXISTING_UINT8", expected)

		if got != expected {
			t.Errorf("GetUint8 should return %v for 'NON_EXISTING_UINT8'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := uint8(123)

		os.Setenv("INVALID_UINT8", "invalid")
		defer os.Unsetenv("INVALID_UINT8")

		got := GetUint8("INVALID_UINT8", expected)

		if got != expected {
			t.Errorf("GetUint8 should return %v for 'INVALID_UINT8'", expected)
		}
	})
}

func TestMustGetUint8(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint8(123)

		os.Setenv("TEST_UINT8", "123")
		defer os.Unsetenv("TEST_UINT8")

		got := MustGetUint8("TEST_UINT8")

		if got != expected {
			t.Errorf("MustGetUint8 should return %v for 'TEST_UINT8'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint8: environment variable 'NON_EXISTING_UINT8' not set"))

		MustGetUint8("NON_EXISTING_UINT8")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint8: invalid environment variable 'TEST_UINT8' has been set: invalid"))

		os.Setenv("INVALID_UINT8", "invalid")
		defer os.Unsetenv("INVALID_UINT8")

		MustGetUint8("INVALID_UINT8")
	})
}

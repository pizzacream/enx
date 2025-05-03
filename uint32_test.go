package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetUint32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint32(123)

		os.Setenv("TEST_UINT32", "123")
		defer os.Unsetenv("TEST_UINT32")

		got := GetUint32("TEST_UINT32", 0)

		if got != expected {
			t.Errorf("GetUint32 should return %v for 'TEST_UINT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := uint32(123)

		got := GetUint32("NON_EXISTING_UINT32", expected)

		if got != expected {
			t.Errorf("GetUint32 should return %v for 'NON_EXISTING_UINT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := uint32(123)

		os.Setenv("INVALID_UINT32", "invalid")
		defer os.Unsetenv("INVALID_UINT32")

		got := GetUint32("INVALID_UINT32", expected)

		if got != expected {
			t.Errorf("GetUint32 should return %v for 'INVALID_UINT32'", expected)
		}
	})
}

func TestMustGetUint32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := uint32(123)

		os.Setenv("TEST_UINT32", "123")
		defer os.Unsetenv("TEST_UINT32")

		got := MustGetUint32("TEST_UINT32")

		if got != expected {
			t.Errorf("MustGetUint32 should return %v for 'TEST_UINT32'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint32: environment variable 'NON_EXISTING_UINT32' not set"))

		MustGetUint32("NON_EXISTING_UINT32")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetUint32: invalid environment variable 'TEST_UINT32' has been set: invalid"))

		os.Setenv("INVALID_UINT32", "invalid")
		defer os.Unsetenv("INVALID_UINT32")

		MustGetUint32("INVALID_UINT32")
	})
}

package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetInt16(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int16(123)

		os.Setenv("TEST_INT16", "123")
		defer os.Unsetenv("TEST_INT16")

		got := GetInt16("TEST_INT16", 0)

		if got != expected {
			t.Errorf("GetInt16 should return %v for 'TEST_INT16'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := int16(123)

		got := GetInt16("NON_EXISTING_INT16", expected)

		if got != expected {
			t.Errorf("GetInt16 should return %v for 'NON_EXISTING_INT16'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := int16(123)

		os.Setenv("INVALID_INT16", "invalid")
		defer os.Unsetenv("INVALID_INT16")

		got := GetInt16("INVALID_INT16", expected)

		if got != expected {
			t.Errorf("GetInt16 should return %v for 'INVALID_INT16'", expected)
		}
	})
}

func TestMustGetInt16(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int16(123)

		os.Setenv("TEST_INT16", "123")
		defer os.Unsetenv("TEST_INT16")

		got := MustGetInt16("TEST_INT16")

		if got != expected {
			t.Errorf("MustGetInt16 should return %v for 'TEST_INT16'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt16: environment variable 'NON_EXISTING_INT16' not set"))

		MustGetInt16("NON_EXISTING_INT16")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt16: invalid environment variable 'TEST_INT16' has been set: invalid"))

		os.Setenv("INVALID_INT16", "invalid")
		defer os.Unsetenv("INVALID_INT16")

		MustGetInt16("INVALID_INT16")
	})
}

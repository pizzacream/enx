package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetInt8(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int8(123)

		os.Setenv("TEST_INT8", "123")
		defer os.Unsetenv("TEST_INT8")

		got := GetInt8("TEST_INT8", 0)

		if got != expected {
			t.Errorf("GetInt8 should return %v for 'TEST_INT8'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := int8(123)

		got := GetInt8("NON_EXISTING_INT8", expected)

		if got != expected {
			t.Errorf("GetInt8 should return %v for 'NON_EXISTING_INT8'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := int8(123)

		os.Setenv("INVALID_INT8", "invalid")
		defer os.Unsetenv("INVALID_INT8")

		got := GetInt8("INVALID_INT8", expected)

		if got != expected {
			t.Errorf("GetInt8 should return %v for 'INVALID_INT8'", expected)
		}
	})
}

func TestMustGetInt8(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int8(123)

		os.Setenv("TEST_INT8", "123")
		defer os.Unsetenv("TEST_INT8")

		got := MustGetInt8("TEST_INT8")

		if got != expected {
			t.Errorf("MustGetInt8 should return %v for 'TEST_INT8'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt8: environment variable 'NON_EXISTING_INT8' not set"))

		MustGetInt8("NON_EXISTING_INT8")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt8: invalid environment variable 'TEST_INT8' has been set: invalid"))

		os.Setenv("INVALID_INT8", "invalid")
		defer os.Unsetenv("INVALID_INT8")

		MustGetInt8("INVALID_INT8")
	})
}

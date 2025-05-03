package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetInt32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int32(123)

		os.Setenv("TEST_INT32", "123")
		defer os.Unsetenv("TEST_INT32")

		got := GetInt32("TEST_INT32", 0)

		if got != expected {
			t.Errorf("GetInt32 should return %v for 'TEST_INT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := int32(123)

		got := GetInt32("NON_EXISTING_INT32", expected)

		if got != expected {
			t.Errorf("GetInt32 should return %v for 'NON_EXISTING_INT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := int32(123)

		os.Setenv("INVALID_INT32", "invalid")
		defer os.Unsetenv("INVALID_INT32")

		got := GetInt32("INVALID_INT32", expected)

		if got != expected {
			t.Errorf("GetInt32 should return %v for 'INVALID_INT32'", expected)
		}
	})
}

func TestMustGetInt32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := int32(123)

		os.Setenv("TEST_INT32", "123")
		defer os.Unsetenv("TEST_INT32")

		got := MustGetInt32("TEST_INT32")

		if got != expected {
			t.Errorf("MustGetInt32 should return %v for 'TEST_INT32'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt32: environment variable 'NON_EXISTING_INT32' not set"))

		MustGetInt32("NON_EXISTING_INT32")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt32: invalid environment variable 'TEST_INT32' has been set: invalid"))

		os.Setenv("INVALID_INT32", "invalid")
		defer os.Unsetenv("INVALID_INT32")

		MustGetInt32("INVALID_INT32")
	})
}

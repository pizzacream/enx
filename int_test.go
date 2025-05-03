package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := 123

		os.Setenv("TEST_INT", "123")
		defer os.Unsetenv("TEST_INT")

		got := GetInt("TEST_INT", 0)

		if got != expected {
			t.Errorf("GetInt should return %v for 'TEST_INT'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := 123

		got := GetInt("NON_EXISTING_INT", 123)

		if got != expected {
			t.Errorf("GetInt should return %v for 'NON_EXISTING_INT'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := 123

		os.Setenv("INVALID_INT", "invalid")
		defer os.Unsetenv("INVALID_INT")

		got := GetInt("INVALID_INT", 123)

		if got != expected {
			t.Errorf("GetInt should return %v for 'INVALID_INT'", expected)
		}
	})
}

func TestMustGetInt(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := 123

		os.Setenv("TEST_INT", "123")
		defer os.Unsetenv("TEST_INT")

		got := MustGetInt("TEST_INT")

		if got != expected {
			t.Errorf("MustGetInt should return %v for 'TEST_INT'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt: environment variable 'NON_EXISTING_INT' not set"))

		MustGetInt("NON_EXISTING_INT")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetInt: invalid environment variable 'TEST_INT' has been set: invalid"))

		os.Setenv("INVALID_INT", "invalid")
		defer os.Unsetenv("INVALID_INT")

		MustGetInt("INVALID_INT")
	})
}

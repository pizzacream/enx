package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := "test"

		os.Setenv("TEST_STRING", expected)
		defer os.Unsetenv("TEST_STRING")

		got := GetString("TEST_STRING", "default")

		if got != expected {
			t.Errorf("GetString should return %v for 'TEST_STRING'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := "default"

		got := GetString("NON_EXISTING_STRING", expected)

		if got != expected {
			t.Errorf("GetString should return %v for 'NON_EXISTING_STRING'", expected)
		}
	})
}

func TestMustGetString(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := "test"

		os.Setenv("TEST_STRING", expected)
		defer os.Unsetenv("TEST_STRING")

		got := MustGetString("TEST_STRING")

		if got != expected {
			t.Errorf("MustGetString should return %v for 'TEST_STRING'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetString: environment variable 'NON_EXISTING_STRING' not set"))

		MustGetString("NON_EXISTING_STRING")
	})
}

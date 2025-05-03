package enx

import (
	"errors"
	"os"
	"testing"
)

func TestGetFloat32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := float32(123.456)

		os.Setenv("TEST_FLOAT32", "123.456")
		defer os.Unsetenv("TEST_FLOAT32")

		got := GetFloat32("TEST_FLOAT32", 0)

		if got != expected {
			t.Errorf("GetFloat32 should return %v for 'TEST_FLOAT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := float32(123.456)

		got := GetFloat32("NON_EXISTING_FLOAT32", 123.456)

		if got != expected {
			t.Errorf("GetFloat32 should return %v for 'NON_EXISTING_FLOAT32'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := float32(123.456)

		os.Setenv("INVALID_FLOAT32", "invalid")
		defer os.Unsetenv("INVALID_FLOAT32")

		got := GetFloat32("INVALID_FLOAT32", 123.456)

		if got != expected {
			t.Errorf("GetFloat32 should return %v for 'INVALID_FLOAT32'", expected)
		}
	})
}

func TestMustGetFloat32(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := float32(123.456)

		os.Setenv("TEST_FLOAT32", "123.456")
		defer os.Unsetenv("TEST_FLOAT32")

		got := MustGetFloat32("TEST_FLOAT32")

		if got != expected {
			t.Errorf("MustGetFloat32 should return %v for 'TEST_FLOAT32'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetFloat32: environment variable 'NON_EXISTING_FLOAT32' not set"))

		MustGetFloat32("NON_EXISTING_FLOAT32")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetFloat32: invalid environment variable 'TEST_FLOAT32' has been set: invalid"))

		os.Setenv("INVALID_FLOAT32", "invalid")
		defer os.Unsetenv("INVALID_FLOAT32")

		MustGetFloat32("INVALID_FLOAT32")
	})
}

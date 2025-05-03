package enx

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestGetDuration(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := time.Second * 10

		os.Setenv("DURATION", "10s")
		defer os.Unsetenv("DURATION")

		got := GetDuration("DURATION", time.Minute)
		if got != expected {
			t.Errorf("GetDuration should return %v for 'DURATION'", expected)
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := time.Minute

		got := GetDuration("DURATION", time.Minute)
		if got != expected {
			t.Errorf("GetDuration should return %v for 'DURATION'", expected)
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := time.Minute

		os.Setenv("DURATION", "invalid")
		defer os.Unsetenv("DURATION")

		got := GetDuration("DURATION", time.Minute)

		if got != expected {
			t.Errorf("GetDuration should return %v for 'DURATION'", expected)
		}
	})
}

func TestMustGetDuration(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := time.Second * 10

		os.Setenv("DURATION", "10s")
		defer os.Unsetenv("DURATION")

		got := MustGetDuration("DURATION")
		if got != expected {
			t.Errorf("MustGetDuration should return %v for 'DURATION'", expected)
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetDuration: environment variable 'DURATION' not set"))

		MustGetDuration("DURATION")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetDuration: invalid environment variable 'DURATION' has been set: invalid"))

		os.Setenv("DURATION", "invalid")
		defer os.Unsetenv("DURATION")

		MustGetDuration("DURATION")
	})
}

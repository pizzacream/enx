package autoload_test

import (
	"os"
	"testing"

	_ "github.com/pizzacream/enx/autoload"
)

func TestLoad(t *testing.T) {
	t.Run("should load environment variables from .env file", func(t *testing.T) {
		os.Setenv("TEST_ENV", "test")
		defer os.Unsetenv("TEST_ENV")

		got := os.Getenv("TEST_ENV")
		if got != "test" {
			t.Errorf("expected 'test', got '%s'", got)
		}
	})
}

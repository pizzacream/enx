package enx

import (
	"os"
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := map[string]string{"key": "value"}

		os.Setenv("TEST_JSON", `{"key":"value"}`)
		defer os.Unsetenv("TEST_JSON")

		got := GetJSON("TEST_JSON", map[string]string{})

		if got["key"] != expected["key"] {
			t.Errorf("GetJSON should return a map with key 'value'")
		}
	})

	t.Run("should return array from environment variable", func(t *testing.T) {
		expected := []string{"value1", "value2"}

		os.Setenv("TEST_JSON", `["value1","value2"]`)
		defer os.Unsetenv("TEST_JSON")

		got := GetJSON("TEST_JSON", []string{})

		if got[0] != expected[0] || got[1] != expected[1] {
			t.Errorf("GetJSON should return an array with values 'value1' and 'value2'")
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := map[string]string{"key": "value"}

		got := GetJSON("NON_EXISTING_JSON", expected)

		if got["key"] != expected["key"] {
			t.Errorf("GetJSON should return the default value for non-existing key")
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := map[string]string{}

		os.Setenv("INVALID_JSON", `{"key":value"}`)
		defer os.Unsetenv("INVALID_JSON")

		got := GetJSON("INVALID_JSON", expected)

		if got["key"] == "value" {
			t.Errorf("GetJSON should return default value for invalid JSON")
		}
	})
}

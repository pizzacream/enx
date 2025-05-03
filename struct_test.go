package enx

import (
	"errors"
	"os"
	"testing"
)

type TestStruct struct {
	Title string `json:"title"`
	Value int    `json:"value"`
}

func TestGetStruct(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := TestStruct{Title: "test", Value: 123}

		os.Setenv("TEST_STRUCT", `{"title":"test","value":123}`)
		defer os.Unsetenv("TEST_STRUCT")

		got := GetStruct("TEST_STRUCT", TestStruct{})

		if got.Title != expected.Title || got.Value != expected.Value {
			t.Errorf("GetStruct should return a TestStruct with title 'test' and value 123")
		}
	})

	t.Run("should return default value if environment variable is not set", func(t *testing.T) {
		expected := TestStruct{Title: "test", Value: 123}

		got := GetStruct("NON_EXISTING_STRUCT", expected)

		if got.Title != expected.Title || got.Value != expected.Value {
			t.Errorf("GetStruct should return a TestStruct with title 'test' and value 123")
		}
	})

	t.Run("should return default value if environment variable is invalid", func(t *testing.T) {
		expected := TestStruct{Title: "test", Value: 123}

		os.Setenv("INVALID_STRUCT", `{"title":"test","value":"not a number"}`)
		defer os.Unsetenv("INVALID_STRUCT")

		got := GetStruct("INVALID_STRUCT", expected)

		if got.Title != expected.Title || got.Value != expected.Value {
			t.Errorf("GetStruct should return a TestStruct with title 'test' and value 123")
		}
	})
}

func TestMustGetStruct(t *testing.T) {
	t.Run("should return value from environment variable", func(t *testing.T) {
		expected := TestStruct{Title: "test", Value: 123}

		os.Setenv("TEST_STRUCT", `{"title":"test","value":123}`)
		defer os.Unsetenv("TEST_STRUCT")

		got := MustGetStruct[TestStruct]("TEST_STRUCT")

		if got.Title != expected.Title || got.Value != expected.Value {
			t.Errorf("MustGetStruct should return a TestStruct with title 'test' and value 123")
		}
	})

	t.Run("should panic if environment variable is not set", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetStruct: environment variable 'NON_EXISTING_STRUCT' not set"))

		MustGetStruct[TestStruct]("NON_EXISTING_STRUCT")
	})

	t.Run("should panic if environment variable is invalid", func(t *testing.T) {
		defer recovery(t, errors.New("MustGetStruct: invalid environment variable 'TEST_STRUCT' has been set: invalid character 'n' looking for beginning of value"))

		os.Setenv("INVALID_STRUCT", `{"title":"test","value":"not a number"}`)
		defer os.Unsetenv("INVALID_STRUCT")

		MustGetStruct[TestStruct]("INVALID_STRUCT")
	})
}

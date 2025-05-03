package enx

import "testing"

func recovery(t *testing.T, err error) {
	if r := recover(); r == nil {
		t.Errorf("expected panic: %v", err)
	}
}

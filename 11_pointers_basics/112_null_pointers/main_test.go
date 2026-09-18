package main

import (
	"strings"
	"testing"
)

func TestGetPtrVal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			err := r.(error)
			expectedMsg := "invalid memory address or nil pointer dereference"
			if !strings.Contains(err.Error(), expectedMsg) {
				t.Errorf("Unexpected panic message. Got: %v, Want: %v", r, expectedMsg)
			}
		}
	}()

	var ptr *int
	// ptr := new(int)
	getPtrVal(ptr)
}

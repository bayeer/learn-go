package main

import (
	"runtime"
	"testing"
)

func TestGetWeakPtr(t *testing.T) {
	ptr := newResource(1, "qwerty")
	got := getWeakPtr(ptr)

	runtime.GC()

	if got.Value() != nil {
		t.Errorf("Want weak ptr nil, got not nil")
	}
}

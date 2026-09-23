package main

import (
	"testing"
	"unsafe"
)

func TestAlignment(t *testing.T) {
	t.Run("sizeof A", func(t *testing.T) {
		assertEqual(t, 16, int(unsafe.Sizeof(A{})))
	})
	t.Run("sizeof B", func(t *testing.T) {
		assertEqual(t, 4, int(unsafe.Sizeof(B{})))
	})
	t.Run("sizeof C", func(t *testing.T) {
		assertEqual(t, 24, int(unsafe.Sizeof(C{})))
	})
	t.Run("sizeof D", func(t *testing.T) {
		assertEqual(t, 16, int(unsafe.Sizeof(D{})))
	})
}

func assertEqual(t *testing.T, expected, got int) {
	t.Helper()

	if got != expected {
		t.Fatalf("expected: %d, got: %d", expected, got)
	}
}

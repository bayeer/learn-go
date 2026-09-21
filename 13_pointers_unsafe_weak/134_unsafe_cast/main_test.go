package main

import (
	"testing"
	"unsafe"
)

func TestToHexString(t *testing.T) {
	t.Run("toHexString", func(t *testing.T) {
		pi := 3.14159
		f := *(*uint64)(unsafe.Pointer(&pi))
		got := toHexString(f)
		want := "0x400921f9f01b866e"
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
	t.Run("toBinString", func(t *testing.T) {
		pi := 3.14159
		f := *(*uint64)(unsafe.Pointer(&pi))
		got := toBinString(f)
		want := "0100000000001001001000011111100111110000000110111000011001101110"
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}

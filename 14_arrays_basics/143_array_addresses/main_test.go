package main

import (
	"fmt"
	"testing"
)

func TestNewArr5(t *testing.T) {
	want := "[5]int"
	got := fmt.Sprintf("%T", NewArr5())
	if want != got {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

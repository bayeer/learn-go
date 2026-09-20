package main

import (
	"testing"
)

var globalPtr *int

func BenchmarkStack(b *testing.B) {
	for b.Loop() {
		x := 42
		_ = x
	}
}

func BenchmarkHeap(b *testing.B) {
	for b.Loop() {
		x := new(int)
		*x = 42
		globalPtr = x // escapes to global var
	}
}

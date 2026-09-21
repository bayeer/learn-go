package main

import (
	"testing"
)

func TestNewEmptySlice(t *testing.T) {
	es := NewEmptySlice()
	got := len(es)
	expected := 0
	if got != expected {
		t.Fatalf("got: %d, expected: %d", got, expected)
	}
}

func TestNewIntSlice(t *testing.T) {
	a := NewIntSlice()
	got := len(a)
	expected := 5
	if got != expected {
		t.Fatalf("got: %d, expected: %d", got, expected)
	}
}

func TestNewStringSlice(t *testing.T) {
	a := NewStringSlice()
	got := len(a)
	expected := 4
	if got != expected {
		t.Fatalf("got: %d, expected: %d", got, expected)
	}
}

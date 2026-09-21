package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytesToString(t *testing.T) {
	want := "asdf"
	b := []byte(want)
	got := bytesToString(b)
	if got != want {
		t.Errorf("want: %q, got: %q\n", want, got)
	}
}

func TestStringToBytes(t *testing.T) {
	s := "asdf"
	want := []byte(s)
	fmt.Printf("want value: %#x\n", want)
	got := stringToBytes(&s)
	fmt.Printf("got value: %#x\n", got)
	if !bytes.Equal(want, got) {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

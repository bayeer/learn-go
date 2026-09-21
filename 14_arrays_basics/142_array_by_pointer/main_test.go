package main

import (
	"testing"
)

func TestModifyByVal(t *testing.T) {
	a := arr5{1, 2, 3, 4, 5}
	want := a
	modifyByValue(want)
	got := a
	if want != got {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

func TestModifyByPtr(t *testing.T) {
	a := arr5{1, 2, 3, 4, 5}
	want := a
	modifyByPtr(&a)
	got := a
	if got == want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

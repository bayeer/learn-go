package main

import "testing"

func TestGetType(t *testing.T) {
	want := "[3]int"

	// array
	a := [...]int{1, 2, 3}
	got := getType(a)
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}

	// slice
	b := []int{1, 2, 3}
	want = "[]int"
	got = getType(b)
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

package main

import (
	"testing"
	"unsafe"
)

func TestGetB(t *testing.T) {
	d := Data{A: 100, B: 200, C: 300}

	want := 200
	got := int(getC(&d))
	if want != got {
		t.Errorf("want: %d, got: %d\n", want, got)
	}
}

func TestGetDataFieldValue(t *testing.T) {
	d := Data{100, 200, 300}

	want := 300
	got := int(
		getDataFieldValue(&d, d.C, unsafe.Offsetof(d.C)),
	)
	if want != got {
		t.Errorf("want: %d, got: %d\n", want, got)
	}
}

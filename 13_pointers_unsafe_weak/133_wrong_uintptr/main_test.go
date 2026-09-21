package main

import (
	"testing"
)

func TestGetValueByUnsafePtr(t *testing.T) {
	x := int64(42)
	got := getValueByUnsafePtr(&x)
	want := int32(4)

	if got != want {
		t.Errorf("got: %d, want: %d\n", got, want)
	}
}

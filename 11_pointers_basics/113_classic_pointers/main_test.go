package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintPersonByRef(t *testing.T) {
	tests := []struct {
		p        Person
		expected string
	}{
		{Person{"Petr", 20}, "name: Petr, age: 20\n"},
		{Person{"Cyrill", 30}, "name: Cyrill, age: 30\n"},
		{Person{"Badma", 35}, "name: Badma, age: 35\n"},
	}

	for _, tt := range tests {
		t.Run(tt.p.name, func(t *testing.T) {
			got := readStdout(func() { printPersonByRef(&tt.p) })
			if got != tt.expected {
				t.Errorf("printPersonByRef() = %s, want %s", got, tt.expected)
			}
		})
	}
}

func readStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w

	f()

	err = w.Close()
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	return buf.String()
}

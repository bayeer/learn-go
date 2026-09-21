package main

import (
	"testing"
)

func TestGetUserHash(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected string
	}{
		{"Vasya", 40, "Vasya#40"},
		{"Kolya", 30, "Kolya#30"},
		{"Marina", 20, "Marina#20"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUser(tt.name, tt.age)

			got := GetUserHash(&u)
			if got != tt.expected {
				t.Errorf("GetUserHash(%q, %d) = %s, want %s", tt.name, tt.age, got, tt.expected)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestUpdateByRef(t *testing.T) {
	tests := []struct {
		user     User
		byVal    bool
		expected string
	}{
		{User{"Petr"}, true, "Petr"},
		{User{"Vasya"}, false, "Bar"},
		{User{"Zoo"}, false, "Bar"},
	}

	for _, tt := range tests {
		t.Run(tt.user.name, func(t *testing.T) {
			var got string
			if tt.byVal {
				updateByVal(tt.user)
				got = tt.user.name

				if got != tt.expected {
					t.Errorf("updateByVal(), want '%s', got '%s'", tt.expected, got)
				}
			} else {
				updateByRef(&tt.user)
				got = tt.user.name

				if got != tt.expected {
					t.Errorf("updateByRef(), want '%s', got '%s'", tt.expected, got)
				}
			}
		})
	}
}

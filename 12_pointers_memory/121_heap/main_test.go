package main

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	initHeapSize := getHeapSize()
	printHeapSize("test:init")

	u := createUser()
	fmt.Println(u.Name)

	hs := getHeapSize()
	printHeapSize("test:after-create-user")

	if hs <= initHeapSize {
		t.Errorf("heap has not changed\n")
	}
}

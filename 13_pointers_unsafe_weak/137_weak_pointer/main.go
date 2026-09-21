package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Resource struct {
	ID   int
	Data string
}

func main() {
	// creating strong pointer.
	ptr := newResource(1, "ASDF")

	// creating weak pointer.
	weakPtr := weak.Make(ptr)

	fmt.Printf("Strong pointer: %+v\n", ptr)
	fmt.Printf("Weak pointer: %+v\n", weakPtr.Value())

	// deleting strong pointer.
	ptr = nil

	runtime.GC()

	if weakPtr.Value() == nil {
		fmt.Printf("Weak pointer was deleted by GC\n")
	}
}

func getWeakPtr(ptr *Resource) weak.Pointer[Resource] {
	return weak.Make(ptr)
}

func newResource(id int, value string) *Resource {
	return &Resource{
		ID:   id,
		Data: value,
	}
}

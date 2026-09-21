package main

import (
	"fmt"
)

func main() {
	numbers := NewIntSlice()

	fmt.Printf("Slice of ints: %v\n", numbers)
	fmt.Printf("Length: %d\n", len(numbers))

	fmt.Println()

	empty := NewEmptySlice()
	fmt.Printf("Empty slice: %v\n", empty)
	fmt.Printf("Length: %d\n", len(empty))

	fmt.Println()

	names := NewStringSlice()
	fmt.Printf("Slice of strings: %v\n", names)
	fmt.Printf("Length: %d\n", len(names))
}

func NewEmptySlice() []int {
	return []int{}
}

func NewIntSlice() []int {
	return []int{1, 2, 3, 4, 5}
}

func NewStringSlice() []string {
	return []string{"Badma", "Erdem", "Zhargal", "Bair"}
}

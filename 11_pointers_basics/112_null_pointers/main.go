package main

import (
	"fmt"
)

func main() {
	var ptr *int
	fmt.Println(ptr)

	// panic
	// fmt.Println(*ptr)

	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Cannot read value of nil pointer")
	}

	n := 42
	ptr = &n
	fmt.Println(*ptr)
}

func getPtrVal(ptr *int) int {
	return *ptr
}

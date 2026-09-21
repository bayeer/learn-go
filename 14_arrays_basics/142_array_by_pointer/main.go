package main

import (
	"fmt"
)

type arr5 [5]int

func main() {
	a := arr5{1, 2, 3, 4, 5}
	fmt.Printf("a before modify: %v\n", a)

	b := modifyByValue(a)
	fmt.Printf("a after modify: %v\n", a)
	fmt.Printf("b after modify: %v\n", b)

	modifyByPtr(&a)
	fmt.Printf("a after modifyByPtr: %v\n", a)
}

func modifyByValue(a arr5) arr5 {
	a[0] = 123
	return a
}

func modifyByPtr(a *arr5) {
	a[0] = 123
}

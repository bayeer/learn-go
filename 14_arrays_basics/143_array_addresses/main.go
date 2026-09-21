package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := NewArr5()

	fmt.Printf("sizeof a: %d\n", unsafe.Sizeof(a))
	fmt.Printf("sizeof a[0]: %d\n", unsafe.Sizeof(a[0]))

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %v, addr: %p\n", i, a[i], &a[i])
	}
}

func NewArr5() [5]int {
	return [...]int{5, 4, 3, 2, 1}
}

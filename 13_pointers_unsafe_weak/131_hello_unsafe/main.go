package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 42
	y := 42.3

	ptr := &x
	fmt.Printf("simple pointer: %T\n", ptr)

	unsafePtr := unsafe.Pointer(&x)
	fmt.Printf("unsafe pointer: %T\n", unsafePtr)

	unsafePtr = unsafe.Pointer(&y)
	fmt.Printf("unsafe pointer: %T\n", unsafePtr)
}

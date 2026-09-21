package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := int64(42)

	fmt.Printf("value by unsafe pointer: %d\n", getValueByUnsafePtr(&x))

	// in the following 2 line, GC might work
	addr := uintptr(unsafe.Pointer(&x))    // dangerous use of uintptr
	ptr2 := (*int64)(unsafe.Pointer(addr)) // can address to wrong place

	fmt.Printf("dangerous value: %d\n", *ptr2) // may fail with panic or wrong data
}

func getValueByUnsafePtr(x *int64) int32 {
	ptr := (*int32)(unsafe.Pointer(x))
	return *(*int32)(ptr)
}

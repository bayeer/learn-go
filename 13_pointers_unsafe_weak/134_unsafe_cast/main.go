package main

import (
	"fmt"
	"unsafe"
)

func main() {
	pi := 3.14159
	bits := *(*uint64)(unsafe.Pointer(&pi))

	fmt.Printf("float64: %f\n", pi)
	fmt.Printf("bits as hex: %q\n", toHexString(bits))
	fmt.Printf("bits as bin: %q\n", toBinString(bits))

	// reverse cast
	pi2 := *(*float64)(unsafe.Pointer(&bits))
	fmt.Printf("restored float64: %f\n", pi2)
}

func toHexString(bits uint64) string {
	return fmt.Sprintf("%#x", bits)
}

func toBinString(bits uint64) string {
	return fmt.Sprintf("%064b", bits)
}

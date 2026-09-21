package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	A int32
	B int64
}

func main() {
	d := Data{A: 10, B: 20}

	fmt.Printf("d: %+v\n", d)

	size := unsafe.Sizeof(d)
	fmt.Printf("sizeof d: %d bytes\n", size)

	align := unsafe.Alignof(d)
	fmt.Printf("alignof d: %d bytes\n", align)

	offset := unsafe.Offsetof(d.B)
	fmt.Printf("offset of d.B: %d bytes\n", offset)
}

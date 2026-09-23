package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int8  // 1 byte
	b int16 // 2 bytes
	c int8  // 1 byte
	d int64 // 8 bytes
} // alignment: 8+8=16

type B struct {
	b int16 // 2 bytes
	a int8  // 1 byte
	c int8  // 1 byte
} // alignment: 2+2=4

type C struct {
	a int8  // 1 byte
	b int64 // 8 bytes
	c int8  // 1 byte
	d int16 // 2 bytes
} // alignment: 8+8+8=24

type D struct {
	b int64 // 8 bytes
	d int16 // 2 bytes
	a int8  // 1 byte
	c int8  // 1 byte
} // alignment: 8+8=16

func main() {
	a := A{}
	fmt.Printf("aligned size of A: %d\n", unsafe.Sizeof(a))
	fmt.Printf("alignment: %d-byte boundaries\n", unsafe.Alignof(a))
	fmt.Printf("offsets, A.a: %d, A.b: %d, A.c: %d, A.d: %d\n\n",
		unsafe.Offsetof(a.a), unsafe.Offsetof(a.b), unsafe.Offsetof(a.c), unsafe.Offsetof(a.d))

	b := B{}
	fmt.Printf("aligned size of B: %d\n", unsafe.Sizeof(b))
	fmt.Printf("alignment: %d-byte boundaries\n", unsafe.Alignof(b))
	fmt.Printf("offsets, B.a: %d, B.b: %d, B.c: %d\n\n",
		unsafe.Offsetof(b.a), unsafe.Offsetof(b.b), unsafe.Offsetof(b.c))

	c := C{}
	fmt.Printf("aligned size of C: %d\n", unsafe.Sizeof(c))
	fmt.Printf("alignment: %d-byte boundaries\n", unsafe.Alignof(c))
	fmt.Printf("offsets, C.a: %d, C.b: %d, C.c: %d, C.d: %d\n\n",
		unsafe.Offsetof(c.a), unsafe.Offsetof(c.b), unsafe.Offsetof(c.c), unsafe.Offsetof(c.d))

	d := D{}
	fmt.Printf("aligned size of C: %d\n", unsafe.Sizeof(d))
	fmt.Printf("alignment: %d-byte boundaries\n", unsafe.Alignof(d))
	fmt.Printf("offsets, D.a: %d, D.b: %d, D.c: %d, D.d: %d\n\n",
		unsafe.Offsetof(d.a), unsafe.Offsetof(d.b), unsafe.Offsetof(d.c), unsafe.Offsetof(d.d))
}

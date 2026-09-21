package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	A int8
	B int64
	C int16
}

type Integer interface {
	~int8 | ~int16 | ~int64
}

func main() {
	d := Data{A: 10, B: 20, C: 30}

	fmt.Printf("sizeof d: %d\n", unsafe.Sizeof(d))
	fmt.Printf("d.A offset: %d\n", unsafe.Offsetof(d.A))
	fmt.Printf("d.B offset: %d\n", unsafe.Offsetof(d.B))
	fmt.Printf("d.C offset: %d\n", unsafe.Offsetof(d.C))

	vB := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&d)) + unsafe.Offsetof(d.B),
	))
	// vB = *(*int64)(unsafe.Add(
	// 	unsafe.Pointer(&d), unsafe.Offsetof(d.B)
	// ))
	fmt.Printf("d.B value by offset: %d\n", *vB)
	fmt.Printf("d.C value by offset: %d\n", getDataFieldValue(&d, d.C, unsafe.Offsetof(d.C)))
}

func getDataFieldValue[T Integer](d *Data, f T, offset uintptr) T {
	basePtr := uintptr(unsafe.Pointer(d))
	ptr := (*T)(unsafe.Pointer(basePtr + uintptr(offset)))
	return *ptr
}

func getC(d *Data) int16 {
	ptr := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(d)) + unsafe.Offsetof(d.C),
	))

	return *ptr
}

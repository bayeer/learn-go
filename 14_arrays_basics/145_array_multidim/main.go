package main

import (
	"fmt"
	"unsafe"
)

func main() {
	mtx := [3][4]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("sizeof mtx: %d bytes\n", unsafe.Sizeof(mtx))
	fmt.Println("elements addresses:\n")

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("mtx[%d][%d] = %2d, addr: %p\n", i, j, mtx[i][j], &mtx[i][j])
		}
	}
}

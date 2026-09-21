package main

import (
	"fmt"
	"time"
)

type LargeArr [100000]int

func main() {
	large := LargeArr{}
	start := time.Now()

	for i := 0; i < 200000; i++ {
		_ = processByVal(large)
	}
	fmt.Printf("By value: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < 200000; i++ {
		_ = processByPtr(&large)
	}
	fmt.Printf("By ptr: %v\n", time.Since(start))
}

func processByVal(arr LargeArr) int {
	return arr[0]
}

func processByPtr(arr *LargeArr) int {
	return arr[0]
}

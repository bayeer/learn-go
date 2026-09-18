package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

var sample []metrics.Sample = []metrics.Sample{
	{Name: "/gc/scan/heap:bytes"},
}

func main() {
	printHeapSize("app:init")
	defer func() { printHeapSize("app:end") }()

	stackExample()
	printHeapSize("app:stack-example")
	runtime.GC()
	printHeapSize("app:stack-example-gc")

	ptr := heapExample()
	fmt.Println(*ptr)
	printHeapSize("app:heap-example")
	runtime.GC()
	printHeapSize("app:heap-example-gc")
}

func stackExample() int {
	x := 42
	return x
}

func heapExample() *int {
	x := 42
	return &x
}

func getHeapSize() int {
	metrics.Read(sample)

	return int(sample[0].Value.Uint64())
}

func printHeapSize(name string) {
	fmt.Printf("heap bytes: %d, '%s'\n", getHeapSize(), name)
}

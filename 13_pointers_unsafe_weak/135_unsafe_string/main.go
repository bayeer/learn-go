package main

import (
	"fmt"
	"runtime/metrics"
	"unsafe"
)

var sample []metrics.Sample = []metrics.Sample{
	{Name: "/gc/scan/heap:bytes"},
}

func main() {
	printHeapSize()
	s := "Hello World!"
	bytes1 := []byte(s)
	printHeapSize()
	fmt.Printf("Bytes from string by copy: %s\n", bytes1)
	printHeapSize()

	// via unsafe
	bytes2 := stringToBytes(&s)
	printHeapSize()
	fmt.Printf("Bytes from string by unsafe: %s\n", bytes2)
	printHeapSize()
}

func stringToBytes(s *string) []byte {
	return unsafe.Slice(unsafe.StringData(*s), len(*s))
}

func bytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func printHeapSize() {
	fmt.Printf("heap size: %d\n", getHeapSize())
}

func getHeapSize() int {
	metrics.Read(sample)

	return int(sample[0].Value.Uint64())
}

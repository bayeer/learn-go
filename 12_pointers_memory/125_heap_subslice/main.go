package main

import (
	"fmt"
	"runtime/metrics"
)

const metricBytes = "/gc/scan/heap:bytes"

var sample = []metrics.Sample{
	{Name: "/gc/scan/heap:bytes"},
}

func main() {
	printHeapSize("app:init")
	defer func() { printHeapSize("app:end") }()

	processBad()
	printHeapSize("app:after-header-bad")

	// runtime.GC()

	// printHeapSize("app:after-gc")
	processGood()
	printHeapSize("app:after-header-good")
}

func getHeaderBad(data []byte) []byte {
	header := data[:100]

	return header
}

func processBad() []byte {
	bigData := make([]byte, 10*1024*1024) // creating 10mb.

	header := getHeaderBad(bigData)

	return header
}

func printHeapSize(name string) {
	fmt.Printf("heap bytes: %d, name: %q\n", getHeapSize(), name)
}

func getHeapSize() int {
	metrics.Read(sample)

	return int(sample[0].Value.Uint64())
}

func getHeaderGood(data []byte) []byte {
	header := make([]byte, 100)

	copy(header, data[:100])

	return header
}

func processGood() []byte {
	bigData := make([]byte, 10*1024*1024)
	header := getHeaderGood(bigData)
	return header
}

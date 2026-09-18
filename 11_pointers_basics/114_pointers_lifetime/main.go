package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

const metricBytes string = "/gc/scan/heap:bytes"

func main() {
	printHeapMetric("app:init", metricBytes)
	var globalPtr *int
	printHeapMetric("app:first-pointer", metricBytes)

	func() {
		temp := 123
		globalPtr = &temp

		fmt.Println(*globalPtr)
	}()

	// temp doesn't exist but value on the heap is still present.
	fmt.Println(*globalPtr)

	printHeapMetric("app:after-anonymous-func", metricBytes)
	// release global pointer so it's value get collected by GC
	globalPtr = nil

	if globalPtr == nil {
		fmt.Println("globalPtr is cleaned")
	}

	printHeapMetric("app:after-globalptr-clean", metricBytes)

	demonstratePtrLifetime()
	runtime.GC()

	printHeapMetric("app:end", metricBytes)
}

func demonstratePtrLifetime() {
	printHeapMetric("app:demo-init", metricBytes)

	data := make([]int, 1)
	data[0] = 42

	ptr := &data[0]

	fmt.Println(*ptr)

	printHeapMetric("app:demo-end", metricBytes)
}

func printHeapMetric(name, key string) {
	sample := []metrics.Sample{
		{Name: key},
	}

	metrics.Read(sample)

	fmt.Printf("metric: '%s': %d, name: '%s'\n",
		key, int(sample[0].Value.Uint64()), name)
}

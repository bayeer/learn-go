package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

type User struct {
	Name string
	Age  int
}

var sample []metrics.Sample = []metrics.Sample{
	{Name: "/gc/scan/heap:bytes"},
}

func main() {
	printHeapSize("app:init")

	defer func() {
		printHeapSize("app:end")
	}()

	user := createUser()
	fmt.Println(user.Name)

	printHeapSize("app:create-user")
	user = nil
	runtime.GC()
	printHeapSize("app:user-nil")
}

func createUser() *User {
	return &User{
		Name: "Alice",
		Age:  18,
	}
}

func getHeapSize() int {
	metrics.Read(sample)

	return int(sample[0].Value.Uint64())
}

func printHeapSize(name string) {
	fmt.Printf("heap bytes: %d, name: %q\n", getHeapSize(), name)
}

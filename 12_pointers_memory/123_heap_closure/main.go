package main

import "fmt"

func main() {
	counter := makeCounter()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func makeCounter() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

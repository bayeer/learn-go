package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println(getType(numbers))

	numbersSlice := []int{1, 2, 3}
	fmt.Println(getType(numbersSlice))
}

func getType(v any) string {
	return fmt.Sprintf("%T", v)
}

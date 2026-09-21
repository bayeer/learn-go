package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := []int{0, 0, 0}

	fmt.Printf("src: %v\n", src)
	fmt.Println()

	fmt.Printf("dst before: %v\n", dst)
	copied := copy(dst, src)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("copied: %d\n", copied)

	dst2 := []int{0, 0, 0, 0, 0}
	fmt.Println()
	fmt.Printf("dst2 before: %v\n", dst2)
	copied2 := copy(dst2, src)
	fmt.Printf("dst2: %v\n", dst2)
	fmt.Printf("copied2: %d\n", copied2)
}

package main

import (
	"fmt"
)

type User struct {
	name string
}

func main() {
	msg := "Hello"
	// variant 1
	// var ptr *string
	// variant 2
	ptr := new(string)
	ptr = &msg

	fmt.Println("value: ", msg)
	fmt.Println("ptr address: ", &msg)
	fmt.Println("ptr: ", ptr)
	fmt.Println("value by ptr: ", *ptr)

	u := User{"Petr"}

	updateByVal(u)
	fmt.Printf("%v\n", u)

	updateByRef(&u)
	fmt.Printf("%v\n", u)
}

func updateByVal(u User) {
	u.name = "Foo"
}

func updateByRef(u *User) {
	u.name = "Bar"
}

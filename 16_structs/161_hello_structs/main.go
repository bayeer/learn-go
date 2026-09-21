package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := NewUser("Vasya", 30)
	fmt.Printf("Name: %q\n", u.Name)
	fmt.Printf("Age: %d\n", u.Age)
}

func GetUserHash(u *User) string {
	return fmt.Sprintf("%s#%d", u.Name, u.Age)
}

func NewUser(n string, a int) User {
	return User{n, a}
}

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	persons := make([]Person, 0)
	persons = append(persons, Person{"Vasya", 27})
	persons = append(persons, Person{"Kolya", 20})
	persons = append(persons, Person{"Anya", 25})

	for _, p := range persons {
		printPersonByRef(&p)
	}
	for _, p := range persons {
		printPersonByVal(p)
	}
}

func printPersonByRef(p *Person) {
	fmt.Printf("name: %s, age: %d\n", (*p).name, (*p).age)
}

func printPersonByVal(p Person) {
	fmt.Printf("name %s, age: %d\n", p.name, p.age)
}

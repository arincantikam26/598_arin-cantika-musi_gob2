package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Airell", age: 23}
	employee1.division = "Curriculum developer"

	// Anonymous struct dengan pengisian field
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee: %+v\n", employee1)
	fmt.Printf("Employee: %+v\n", employee2)

}

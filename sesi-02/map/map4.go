package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("Before deleting", person)

	delete(person, "age")

	fmt.Println("After deleteting: ", person)

}

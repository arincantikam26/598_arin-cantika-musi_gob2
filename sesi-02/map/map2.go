package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address", person["address"])

}

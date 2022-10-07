package main

import "fmt"

func main() {

	var person = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}

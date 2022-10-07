package main

import "fmt"

func main() {
	var person map[string]string //deklarasi

	person = map[string]string{}

	person["name"] = "Airell"

	person["age"] = "23"

	person["address"] = "Jalan Sudirman"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address", person["address"])

}

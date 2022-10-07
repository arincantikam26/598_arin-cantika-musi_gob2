package main

import "fmt"

func main() {
	greet("Airell", "Jalan Sudirman")

}

func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)

}

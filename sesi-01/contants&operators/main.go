package main

import (
	"fmt"
	"strings"
)

func main() {

	// 01 - constant
	const full_name string = "Airell Jordan"

	fmt.Printf("Hello %s", full_name)
	fmt.Println(strings.Repeat("-", 30))

	// 02 - Operators (Relational Operators) 1
	var value1 = 2 + 2*3
	var value2 = (2 + 2) * 3
	fmt.Println("Value 1 = ", value1)
	fmt.Println("Value 2 = ", value2)
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Operators (Relational Operators) 2
	var fsCondt bool = 2 < 3
	var sCondt bool = "joey" == "Joey"
	var tCondt bool = 10 != 2.3
	var frCondt bool = 11 <= 11

	fmt.Println("first condition: ", frCondt)
	fmt.Println("second condition: ", sCondt)
	fmt.Println("third condition: ", tCondt)
	fmt.Println("four condition: ", fsCondt)
	fmt.Println(strings.Repeat("-", 30))

	// 04 - Operators (Relational Operators) 3
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("wrong \t\t(%t) \n", wrongReverse)
	fmt.Println(strings.Repeat("-", 30))
}

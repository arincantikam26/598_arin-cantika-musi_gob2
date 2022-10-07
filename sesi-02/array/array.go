package main

import (
	"fmt"
	"strings"
)

func main() {
	// 01 - array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var str = [3]string{"Airell", "Nanda", "Nailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", str)

	fmt.Println(strings.Repeat("-", 30))

	// 02 - Array(Modify Element Through Index)
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Array(Loop through elements)

	// cara pertama
	for i, v := range fruits {
		fmt.Printf("Index : %d, value: %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))

	// cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index : %d, value: %s\n", i, fruits[i])
	}

	fmt.Println(strings.Repeat("-", 30))

	// 04 - Array(Multidimensional Array)

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}

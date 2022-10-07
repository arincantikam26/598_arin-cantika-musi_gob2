package main

import (
	"fmt"
	"strings"
)

func main() {
	// 01 - slice
	var fruits = []string{"apple", "banana", "mango"}

	_ = fruits

	// 02 - Slice (make function)
	var fruits2 = make([]string, 3)

	_ = fruits2

	fmt.Printf("%#v", fruits)
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Slice (append function)
	var fruits3 = make([]string, 3)

	fruits3 = append(fruits3, "apple", "banana", "mango")

	fmt.Printf("%#v", fruits3)
	fmt.Println(strings.Repeat("-", 30))

	// 04 - Slice (append function with ellipsis)
	var ft1 = []string{"apple", "banana", "mango"}
	var ft2 = []string{"durian", "pineapple", "starfruit"}

	ft1 = append(ft1, ft2...)

	fmt.Printf("%#v", ft1)

	// 05 - Slice (copy function)
	var frt1 = []string{"apple", "banana", "mango"}
	var frt2 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(frt1, frt2)

	fmt.Println("Fruits1 =>", frt1)
	fmt.Println("Fruits1 =>", frt2)
	fmt.Println("Copied elements =>", nn)
	fmt.Println(strings.Repeat("-", 30))

	// 06 - Slice (slicing)
	var fts1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fts2 = fts1[1:4]
	fmt.Printf("%#v\n", fts2)

	var fts3 = fts1[0:]
	fmt.Printf("%#v\n", fts3)

	var fts4 = fts1[:3]
	fmt.Printf("%#v\n", fts4)

	var fts5 = fts1[:]
	fmt.Printf("%#v\n", fts5)
	fmt.Println(strings.Repeat("-", 30))

	// 07 - Slice (Combining slicing and append)
	var ftu1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	ftu1 = append(ftu1[:3], "rambutan")

	fmt.Printf("%#v\n", ftu1)
	fmt.Println(strings.Repeat("-", 30))

	// 08 - Slice (Backing array)
	var ftu2 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var ftu3 = ftu2[2:4]
	ftu3[0] = "rambutan"

	fmt.Println("fruits1 => ", ftu2)
	fmt.Println("fruits2 => ", ftu3)

	// 09 - (Cap Function)
	var ftsd1 = []string{"apple", "mango", "durian", "banana"}
	fmt.Println("Fruits1 cap: ", cap(ftsd1))
	fmt.Println("Fruits1 len: ", len(ftsd1))

	fmt.Println(strings.Repeat("#", 20))

	var ftsd2 = ftsd1[0:3]

	fmt.Println("Fruits 2 cap: ", cap(ftsd2))
	fmt.Println("Fruits 2 len: ", len(ftsd2))

	fmt.Println(strings.Repeat("#", 20))

	var ftsd3 = ftsd1[1:]
	fmt.Println("Fruits 3 cap: ", cap(ftsd3))
	fmt.Println("Fruits 3 len: ", len(ftsd3))

	fmt.Println(strings.Repeat("-", 30))

	// 10 - (Creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("Cars: ", cars)
	fmt.Println("newCars: ", newCars)

}

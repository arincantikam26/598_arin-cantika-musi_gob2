package main

import (
	"fmt"
	"strings"
)

func main() {

	// 01 - Loopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println(strings.Repeat("-", 30))

	// 02 - Loopings (second way of looping)
	var x = 0

	for x < 3 {
		fmt.Println("Angka", x)
		x++
	}
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Loopings (third way of looping)
	var c = 0

	for {
		fmt.Println("Angka", c)
		c++
		if c == 3 {
			break
		}
	}
	fmt.Println(strings.Repeat("-", 30))

	// 04 - Loopings (break and continue keywords)
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
	fmt.Println(strings.Repeat("-", 30))

	// 05 - Nested Looping
	for a := 0; a < 5; a++ {
		for b := a; b < 5; b++ {
			fmt.Print(b, " ")
		}
		fmt.Print()
	}
	fmt.Println(strings.Repeat("-", 30))

	// 06 - Loopings (Label)
	for d := 0; d < 3; d++ {
		fmt.Println("Perulangan ke - ", d+1)
		for e := 0; e < 3; e++ {
			if d == 2 {
				break
			}
			fmt.Print(e, " ")
		}
		fmt.Print("\n")
	}

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 01 - Switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	fmt.Println(strings.Repeat("-", 30))

	// 02 - Switch with relational operators
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")

	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Switch (fallthrough keyword)
	var score3 = 6

	switch {
	case score3 == 8:
		fmt.Println("perfect")
	case (score3 < 8) && (score3 > 3):
		fmt.Println("not bad")
		fallthrough
	case score3 < 5:
		fmt.Println("it is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
	fmt.Println(strings.Repeat("-", 30))
}

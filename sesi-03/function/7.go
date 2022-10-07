package main

import "fmt"

func main() {
	numberlists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberlists...)

	fmt.Println("Result: ", result)

}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total

}

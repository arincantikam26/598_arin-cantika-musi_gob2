package main

import "fmt"

func main() {
	type second = uint
	var hour second = 2600
	fmt.Printf("hour type : %T\n", hour)
}

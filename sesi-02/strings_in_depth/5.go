package main

import "fmt"

func main() {
	str := "manca"

	for i, s := range str {
		fmt.Printf("Index => %d, string => %s\n", i, string(s))
	}

}

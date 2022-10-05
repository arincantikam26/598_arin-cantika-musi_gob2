package main

import (
	"fmt"
	"strings"
)

func main() {
	// 01-Number (integer) 1
	var first = 89
	var second = -12

	fmt.Printf("Tipe data first : %T\n", first)
	fmt.Printf("Tipe data second : %T\n", second)
	fmt.Println(strings.Repeat("-", 30))

	// 02-Number (integer) 2
	var first1 uint = 89
	var second1 int8 = -12

	fmt.Printf("Tipe data first : %T\n", first1)
	fmt.Printf("Bilangan second : %T\n", second1)
	fmt.Println(strings.Repeat("-", 30))

	// 03-Number (float)
	var decimalNumber float32 = 3.63

	fmt.Printf("Decimal Number : %f\n", decimalNumber)
	fmt.Printf("Decimal Number : %.3f\n", decimalNumber)
	fmt.Println(strings.Repeat("-", 30))

	// 04-Bool
	var condition bool = true

	fmt.Printf("is it permitted?  %t \n", condition)
	fmt.Println(strings.Repeat("-", 30))

	// 05-String 1
	var msg1 string = "Halo"

	fmt.Print(msg1)
	fmt.Println(strings.Repeat("-", 30))

	// 06-String 2
	var msg2 = `Selamat datang di "Hacktiv8".
	salam kenal :).
	Mari belajar "Scalable Web Service With Go".`

	fmt.Println(msg2)
	fmt.Println(strings.Repeat("-", 30))
}

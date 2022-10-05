package main

import (
	"fmt"
	"strings"
)

func main() {

	// 01 - Variable with data type

	var name string = "Airell"
	var age int = 23

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umunya ==>", age)
	fmt.Println(strings.Repeat("-", 30))

	// 02 - Variable without data type

	var name1 = "Airell"
	var age1 = 23

	fmt.Printf("%T, %T \n", name1, age1)
	fmt.Println(strings.Repeat("-", 30))

	// 03 - Variable without data type(Short Declaration)

	name2 := "Airell"
	age2 := 23

	fmt.Printf("%T, %T \n", name2, age2)
	fmt.Println(strings.Repeat("-", 30))

	// 04 - Multiple variable declarations 1
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)
	fmt.Println(strings.Repeat("-", 30))

	// 05 - Multiple variable declarations 2
	var st1, st2, st3 = "satu", "dua", "tiga"

	f, s, t := "1", 2, "3"

	fmt.Println(st1, st2, st3)
	fmt.Println(f, s, t)
	fmt.Println(strings.Repeat("-", 30))

	// 06 - Underscore variable
	var firstVariable string

	var name3, age3, address3 = "Airell", 23, "Jalan sudirman"

	_, _, _, _ = firstVariable, name3, age3, address3

	// 07 - fmt.Printf Usage 1

	first1, second1 := 1, "2"

	fmt.Printf("Tipe data variable first adalah %T \n", first1)
	fmt.Printf("Tipe data variable second adalah %T \n", second1)
	fmt.Println(strings.Repeat("-", 30))

	// 08 - fmt.Printf Usage 2
	var name4 = "Airell"
	var age4 = 23
	var address4 = "Jalan Sudirman"

	fmt.Printf("Halo nama ku %s, umurku adalah %d, dan aku tinggal di %s", name4, age4, address4)
	fmt.Println(strings.Repeat("-", 30))

}

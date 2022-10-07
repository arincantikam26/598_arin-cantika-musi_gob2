package main

import (
	"fmt"
	"reflect"
	"strings"
)

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	//identifying data type & value
	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("tipe variabel : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel : ", reflectValue.Int())
	}
	fmt.Println(strings.Repeat("-", 30))

	//accessing value using interface method
	fmt.Println("tipe variable : ", reflectValue.Type())
	fmt.Println("nilai variable : ", reflectValue.Interface())
	fmt.Println(strings.Repeat("-", 30))

	nilai := reflectValue.Interface().(int)
	fmt.Println(nilai)
	fmt.Println(strings.Repeat("-", 30))

	//identifying method information
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("Nama : ", s1.Name)

	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})
	fmt.Println("nama : ", s1.Name)

}

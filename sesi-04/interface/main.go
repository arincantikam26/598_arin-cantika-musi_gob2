package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.widht
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.widht)
}

//interface 3
func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

//type assertion
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	//interface 1
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{widht: 3, height: 2}

	fmt.Printf("Type of c1: %T \n", c1)
	fmt.Printf("Type of r1: %T \n", r1)
	fmt.Println(strings.Repeat("-", 30))

	//interface 2
	fmt.Println("Circle area", c1.area())
	fmt.Println("Circle perimeter", c1.perimeter())
	fmt.Println(strings.Repeat("-", 30))

	//interface 3
	print("Rectangle", c1)
	print("Circle", r1)
	fmt.Println(strings.Repeat("-", 30))

	//type assertion 1
	c1.(circle).volume()
	fmt.Println(strings.Repeat("-", 30))

	//type assertion 2
	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value : %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
		fmt.Println(strings.Repeat("-", 30))
	}

}

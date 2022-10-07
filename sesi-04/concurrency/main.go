package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	//1
	go goroutine()
	fmt.Println(strings.Repeat("-", 30))

	//Asynchronous process #1

	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)
	fmt.Println("No. of Goroutine", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")
	fmt.Println(strings.Repeat("-", 30))

}

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

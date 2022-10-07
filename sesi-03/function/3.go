package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Airell", "Jordan"}
	var printMsg = greet("Heii", names)

	fmt.Println(printMsg)

}

func greet(msg string, name []string) string {
	var joinstr = strings.Join(name, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinstr)

	return result
}

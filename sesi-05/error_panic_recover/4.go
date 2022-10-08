package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()

}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured: ", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	return "Valid password", nil
}

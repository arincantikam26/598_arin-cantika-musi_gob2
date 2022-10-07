package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args

	showData(input[1])

}

type Biodata struct {
	Name       string
	Address    string
	Profession string
	Reason     string
}

func showData(choice string) {
	oneData := Biodata{
		Name:       "Arin Cantika Musi",
		Address:    "Padang",
		Profession: "Student",
		Reason:     "Because golang is a language thas is easy to understand",
	}

	twoData := Biodata{
		Name:       "Vini Jumatul Fitri",
		Address:    "Riau",
		Profession: "Student",
		Reason:     "Because I want to depth learn in golang",
	}
	if choice == "1" {
		fmt.Printf("Absen 1 \n Name : %s \n Address : %s \n Profession : %s \n Reason : %s", oneData.Name, oneData.Address, oneData.Profession, oneData.Reason)
	} else if choice == "2" {
		fmt.Printf("Absen 2 \n Name : %s \n Address : %s \n Profession : %s \n Reason : %s", twoData.Name, twoData.Address, twoData.Profession, twoData.Reason)
	} else {
		fmt.Println("Student doesn't exist!")
	}

}

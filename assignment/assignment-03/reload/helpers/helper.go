package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomNum(min, max int) int {
	return min + rand.Intn(max-min)
}

func WaterUnit(number int) string {
	desc := fmt.Sprintln(number, " m")
	return desc
}

func WindUnit(number int) string {
	desc := fmt.Sprintln(number, " m/s")
	return desc
}

func WaterStatus(number int) string {
	data := ""
	if number < 5 {
		data = "Aman"
	} else if number > 5 && number < 8 {
		data = "Siaga"
	} else {
		data = "Bahaya"
	}
	return data

}

func WindStatus(number int) string {
	data := ""
	if number < 5 {
		data = "Aman"
	} else if number > 5 && number < 8 {
		data = "Siaga"
	} else {
		data = "Bahaya"
	}
	return data

}

// Open file json and read data
func GetData() Data {
	fileBytes, err := ioutil.ReadFile("./jsons/data.json")

	if err != nil {
		panic(err)
	}

	var data Data

	err = json.Unmarshal(fileBytes, &data)

	if err != nil {
		panic(err)
	}

	return data
}

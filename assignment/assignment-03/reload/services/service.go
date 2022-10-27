package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reload/helpers"
)

func UpdateJSON(disaster helpers.Data) {
	// Update (write) the JSON File

	const FilePath = "./jsons/data.json"

	disaster.Status.Water = helpers.RandomNum(1, 100)
	disaster.Status.Wind = helpers.RandomNum(1, 100)

	disasterByte, err := json.Marshal(disaster)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(FilePath, disasterByte, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", disaster)
}

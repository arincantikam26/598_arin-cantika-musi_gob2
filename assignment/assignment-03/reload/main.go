package main

import (
	"net/http"
	"reload/controllers"
	"reload/helpers"
	"reload/services"
	"time"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", controllers.DataController)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

}

func autoUpdate() {
	data := helpers.GetData()
	for range time.Tick(15 * time.Second) {
		services.UpdateJSON(data)
	}
}

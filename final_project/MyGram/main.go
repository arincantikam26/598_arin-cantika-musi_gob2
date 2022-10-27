package main

import (
	"MyGram/databases"
	"MyGram/routers"
)

func main() {
	databases.StartDB()
	r := routers.StartApp()
	r.Run(":3000")

}

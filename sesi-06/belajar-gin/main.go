package main

import "belajar-gin/routes"

func main() {
	var PORT = ":8080"

	routes.StartServer().Run(PORT)

}

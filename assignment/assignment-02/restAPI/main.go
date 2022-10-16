package main

import "restAPI/routes"

func main() {
	PORT := ":3000"

	routes.StartServer().Run(PORT)

}

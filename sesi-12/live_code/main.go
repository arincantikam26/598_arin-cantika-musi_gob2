package main

import "live_code/routes"

func main() {

	PORT := ":3000"
	routes.StartServer().Run(PORT)

}

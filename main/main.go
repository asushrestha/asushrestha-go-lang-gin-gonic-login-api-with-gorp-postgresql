package main

import "gin-login/mappings"

func main() {

	mappings.CreateUrlMappings()

	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")

}

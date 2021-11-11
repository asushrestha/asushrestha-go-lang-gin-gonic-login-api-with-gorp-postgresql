package main

import (
	"gin-login/mappings"
	"os"
)

func main() {

	mappings.CreateUrlMappings()

	// Listen and server on 0.0.0.0:8080
	port, err := os.Getenv("PORT")
	if err != nil {
		port = 8000
	}
	mappings.Router.Run(port)

}

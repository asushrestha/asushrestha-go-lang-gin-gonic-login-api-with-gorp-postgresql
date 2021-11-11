package main

import (
	"gin-login/mappings"
	"os"
)

func main() {

	mappings.CreateUrlMappings()

	// Listen and server on 0.0.0.0:8080
	port := os.Getenv("PORT")
	mappings.Router.Run(":" + port)

}

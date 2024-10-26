package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	app.routes(router)

	// Start the server
	err := router.Run(fmt.Sprintf(":%s", webPort))
	if err != nil {
		log.Panic(err)
	}
}

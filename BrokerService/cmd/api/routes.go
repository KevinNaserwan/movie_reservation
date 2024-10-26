package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// routes method to define your routes
func (app *Config) routes(router *gin.Engine) {
	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.GET("/ping", app.heartbeat)
	router.POST("/", app.GinBroker) // Using the Gin-compatible handler
	router.POST("/handle", app.HandleSubmission)
}

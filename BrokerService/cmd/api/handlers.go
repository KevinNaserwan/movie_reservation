package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Action string `json:"action"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// heartbeat handler for the /ping route
func (app *Config) heartbeat(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (app *Config) GinBroker(c *gin.Context) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	c.JSON(http.StatusOK, payload)
}

func (app *Config) HandleSubmission(c *gin.Context) {
	var requestPayload RequestPayload

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	switch requestPayload.Action {
	default:
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "unknown action",
		})
	}
}

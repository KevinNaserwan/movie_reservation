package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kevinnaserwan/MovieService/repository"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	// Define the routes
	router.GET("/movies", server.listMovies)
	router.GET("/movies/:id", server.getMovie)
	router.POST("/movies", server.createMovie)
	router.PUT("/movies/:id", server.updateMovie)
	router.DELETE("/movies/:id", server.deleteMovie)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

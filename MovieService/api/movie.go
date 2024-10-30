package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/kevinnaserwan/MovieService/repository"
)

type createMovieRequest struct {
	Title       string             `json:"title" binding:"required"`
	Genre       string             `json:"genre" binding:"required"`
	Duration    int32              `json:"duration" binding:"required"`
	ReleaseDate pgtype.Timestamptz `json:"release_date" binding:"required"`
	Rating      float32            `json:"rating" binding:"required"`
	Description string             `json:"description"`
}

func (server *Server) createMovie(ctx *gin.Context) {
	var req createMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMovieParams{
		Title:       req.Title,
		Genre:       req.Genre,
		Duration:    req.Duration,
		ReleaseDate: req.ReleaseDate,
		Rating:      req.Rating,
		Description: req.Description,
	}

	movie, err := server.store.CreateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

type getMovieRequest struct {
	ID int64 `uri:"movie_id" binding:"required,min=1"`
}

func (server *Server) getMovie(ctx *gin.Context) {
	var req getMovieRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movie, err := server.store.GetMovie(ctx, int32(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

type listMoviesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMovies(ctx *gin.Context) {
	var req listMoviesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMoviesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	movies, err := server.store.ListMovies(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

type updateMovieRequest struct {
	Title       string             `json:"title" binding:"required"`
	Genre       string             `json:"genre" binding:"required"`
	Duration    int32              `json:"duration" binding:"required"`
	ReleaseDate pgtype.Timestamptz `json:"release_date" binding:"required"`
	Rating      float32            `json:"rating" binding:"required"`
	Description string             `json:"description"`
}

func (server *Server) updateMovie(ctx *gin.Context) {
	var req updateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMovieParams{
		MovieID:     int32(id),
		Title:       req.Title,
		Genre:       req.Genre,
		Duration:    req.Duration,
		ReleaseDate: req.ReleaseDate,
		Rating:      req.Rating,
		Description: req.Description,
	}

	movie, err := server.store.UpdateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (server *Server) deleteMovie(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteMovie(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movie.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  title, genre, duration, release_date, rating, description
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING movie_id, title, genre, duration, release_date, rating, description
`

type CreateMovieParams struct {
	Title       pgtype.Text
	Genre       pgtype.Text
	Duration    pgtype.Int4
	ReleaseDate pgtype.Date
	Rating      pgtype.Numeric
	Description pgtype.Text
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, createMovie,
		arg.Title,
		arg.Genre,
		arg.Duration,
		arg.ReleaseDate,
		arg.Rating,
		arg.Description,
	)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.Title,
		&i.Genre,
		&i.Duration,
		&i.ReleaseDate,
		&i.Rating,
		&i.Description,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movie_id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, movieID int32) error {
	_, err := q.db.Exec(ctx, deleteMovie, movieID)
	return err
}

const getMovie = `-- name: GetMovie :one
SELECT movie_id, title, genre, duration, release_date, rating, description FROM movies
WHERE movie_id = $1
LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, movieID int32) (Movie, error) {
	row := q.db.QueryRow(ctx, getMovie, movieID)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.Title,
		&i.Genre,
		&i.Duration,
		&i.ReleaseDate,
		&i.Rating,
		&i.Description,
	)
	return i, err
}

const getMoviesByGenre = `-- name: GetMoviesByGenre :many
SELECT movie_id, title, genre, duration, release_date, rating, description FROM movies
WHERE genre = $1
ORDER BY movie_id
LIMIT $2
OFFSET $3
`

type GetMoviesByGenreParams struct {
	Genre  pgtype.Text
	Limit  int32
	Offset int32
}

func (q *Queries) GetMoviesByGenre(ctx context.Context, arg GetMoviesByGenreParams) ([]Movie, error) {
	rows, err := q.db.Query(ctx, getMoviesByGenre, arg.Genre, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.MovieID,
			&i.Title,
			&i.Genre,
			&i.Duration,
			&i.ReleaseDate,
			&i.Rating,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMovies = `-- name: ListMovies :many
SELECT movie_id, title, genre, duration, release_date, rating, description FROM movies
ORDER BY movie_id
LIMIT $1
OFFSET $2
`

type ListMoviesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error) {
	rows, err := q.db.Query(ctx, listMovies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.MovieID,
			&i.Title,
			&i.Genre,
			&i.Duration,
			&i.ReleaseDate,
			&i.Rating,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE movies
SET title = $2, genre = $3, duration = $4, release_date = $5, rating = $6, description = $7
WHERE movie_id = $1
RETURNING movie_id, title, genre, duration, release_date, rating, description
`

type UpdateMovieParams struct {
	MovieID     int32
	Title       pgtype.Text
	Genre       pgtype.Text
	Duration    pgtype.Int4
	ReleaseDate pgtype.Date
	Rating      pgtype.Numeric
	Description pgtype.Text
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, updateMovie,
		arg.MovieID,
		arg.Title,
		arg.Genre,
		arg.Duration,
		arg.ReleaseDate,
		arg.Rating,
		arg.Description,
	)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.Title,
		&i.Genre,
		&i.Duration,
		&i.ReleaseDate,
		&i.Rating,
		&i.Description,
	)
	return i, err
}

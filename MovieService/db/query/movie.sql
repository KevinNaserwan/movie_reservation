-- name: CreateMovie :one
INSERT INTO movies (
  title, genre, duration, release_date, rating, description
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetMovie :one
SELECT * FROM movies
WHERE movie_id = $1
LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
ORDER BY movie_id
LIMIT $1
OFFSET $2;

-- name: UpdateMovie :one
UPDATE movies
SET title = $2, genre = $3, duration = $4, release_date = $5, rating = $6, description = $7
WHERE movie_id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movie_id = $1;

-- name: GetMoviesByGenre :many
SELECT * FROM movies
WHERE genre = $1
ORDER BY movie_id
LIMIT $2
OFFSET $3;

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	MovieID     int32     `json:"movie_id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	Duration    int32     `json:"duration"`
	ReleaseDate pgtype.Timestamptz `json:"release_date"`
	Rating      float32   `json:"rating"`
	Description string    `json:"description"`
}
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	MovieID     int32
	Title       pgtype.Text
	Genre       pgtype.Text
	Duration    pgtype.Int4
	ReleaseDate pgtype.Date
	Rating      pgtype.Numeric
	Description pgtype.Text
}

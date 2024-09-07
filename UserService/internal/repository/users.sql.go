// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, email, role, otp, verified_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, username, email, role, otp, verified_at, created_at, updated_at
`

type CreateUserParams struct {
	ID         pgtype.UUID
	Username   string
	Email      string
	Role       string
	Otp        pgtype.Int4
	VerifiedAt pgtype.Timestamp
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Role,
		arg.Otp,
		arg.VerifiedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Otp,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, role, otp, verified_at, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Otp,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, role, otp, verified_at, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Otp,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $2, email = $3, role = $4, otp = $5, verified_at = $6, updated_at = now()
WHERE id = $1
`

type UpdateUserParams struct {
	ID         pgtype.UUID
	Username   string
	Email      string
	Role       string
	Otp        pgtype.Int4
	VerifiedAt pgtype.Timestamp
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Role,
		arg.Otp,
		arg.VerifiedAt,
	)
	return err
}
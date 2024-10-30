package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

// store provides all functions to execute db queries and transactions
type PGXStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *PGXStore {
	return &PGXStore{
		Queries: New(db),
		db:      db,
	}
}

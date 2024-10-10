package queries

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	*Queries
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		Queries: New(db),
		db:      db,
	}
}

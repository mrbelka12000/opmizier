package repository

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	defaultDBLoad = 10_000_000
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) List(ctx context.Context, query string, args []any) error {

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("repository.List: %w", err)
	}
	defer rows.Close()

	return rows.Err()
}

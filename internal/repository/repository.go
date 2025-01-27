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

func (r *Repo) List(ctx context.Context, query string) error {

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	err = rows.Close()
	if err != nil {
		return fmt.Errorf("rows close: %w", err)
	}

	return rows.Err()
}

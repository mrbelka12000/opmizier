package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Connect
func Connect() (*sql.DB, error) {
	pgURL, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		return nil, fmt.Errorf("environment variable 'POSTGRES_URL' is not set")
	}

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	return db, nil
}

package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

// Client wraps sql.DB
type Client struct {
	DB *sql.DB
}

func NewClient(dsn string) (*Client, error) {
	if dsn == "" {
		return nil, errors.New("DSN cannot be empty")
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database : %w", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		err = fmt.Errorf("ping database: %w", pingErr)
		return nil, err
	}

	return &Client{DB: db}, nil
}

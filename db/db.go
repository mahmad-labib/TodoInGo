package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"time"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Client wraps sql.DB
type Client struct {
	DB *sql.DB
}

func NewClient(dsn string) (*Client, []string, error) {
	if dsn == "" {
		return nil, nil, errors.New("DSN cannot be empty")
	}
	var msg []string
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, msg, fmt.Errorf("open database : %w", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		err = fmt.Errorf("ping database: %w", pingErr)
		return nil, msg, err
	} else {
		msg = append(msg, "connection established succesfully")
		// not sure should migration start with connection or with specific command only!!
		// result, migErr := MigrateUp(db)
		// msg = append(msg, result)
		// err = errors.Join(err, migErr)
	}

	return &Client{DB: db}, msg, nil
}

func MigrateUp(db *sql.DB) (string, error) {
	var msg string
	migrationFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	provider, err := goose.NewProvider("sqlite3", db, migrationFS)
	defer cancel()
	if err != nil {
		return msg, fmt.Errorf("initialize migration provider: %w", err)
	}
	result, err := provider.Up(ctx)
	for _, rs := range result {
		msg += rs.String()
	}
	if err != nil {
		return msg, fmt.Errorf("in migration up: %w", err)
	}
	return msg, nil
}

func MigrateDown(db *sql.DB) (string, error) {
	migrationFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return "", err
	}
	var msg string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	provider, err := goose.NewProvider("sqlite3", db, migrationFS)
	if err != nil {
		return msg, fmt.Errorf("initialize migration provider: %w", err)
	}
	result, err := provider.Down(ctx)
	if result != nil {
		msg += result.String()
	}
	if err != nil {
		return msg, fmt.Errorf("in migration down: %w", err)
	}
	return msg, nil
}

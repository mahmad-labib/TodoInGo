package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"time"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func MigrateUp(db *sql.DB) (string, error) {
	var msg string

	migrationFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	provider, err := goose.NewProvider("sqlite3", db, migrationFS)

	if err != nil {
		return msg, fmt.Errorf("initialize migration provider: %w", err)
	}

	result, err := provider.Up(ctx)

	if err != nil {
		return msg, fmt.Errorf("in migration up: %w", err)
	}
	for _, rs := range result {
		msg += rs.String()
	}

	return msg, nil
}

func MigrateDown(db *sql.DB) (string, error) {
	var msg string

	migrationFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return "", err
	}

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

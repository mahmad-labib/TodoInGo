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

func InitDB() (*sql.DB, []string, error) {
	var msg []string
	db, err := sql.Open("sqlite", "todo.db")
	if err != nil {
		return db, msg, fmt.Errorf("open database : %w", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		err = fmt.Errorf("ping database: %w", pingErr)
	} else {
		msg = append(msg, "connection established succesfully")
		result, migErr := MigrateUp(db)
		msg = append(msg, result)
		err = errors.Join(err, migErr)
	}

	return db, msg, err
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
	if result != nil {
		msg += "create_todo_table"
	}
	if err != nil {
		return msg, fmt.Errorf("in migration up: %w", err)
	}
	return msg, err
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
		msg += "drop_todo_table"
	}
	if err != nil {
		return msg, fmt.Errorf("in migration down: %w", err)
	}
	return msg, err
}

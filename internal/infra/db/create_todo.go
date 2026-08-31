package db

import (
	"context"
	"database/sql"
	"fmt"
)

func TodoCreate(ctx context.Context, db *sql.DB, title, priority, status string) (int64, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err := fmt.Errorf("failed to begin tx: %w", err)
		return 0, err
	}

	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, "INSERT INTO todo (title, status, priority) VALUES (?, ?, ?)", title, status, priority)
	if err != nil {
		err := fmt.Errorf("failed to add new todo with err: %w", err)
		return 0, err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		err := fmt.Errorf("failed to add new todo with err: %w", err)
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		err := fmt.Errorf("failed to commit tx: %w", err)
		return 0, err
	}

	return ID, nil
}

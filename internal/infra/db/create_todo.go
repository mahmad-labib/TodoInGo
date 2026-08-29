package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"slices"
)

type Priority struct {
	Name string
}

func (P Priority) validate() error {
	if P.Name == "" {
		err := fmt.Errorf("priority value should be only 'high', 'normal' or 'low' got nil")
		return err
	}
	priorities := []string{"high", "normal", "low"}
	if exist := slices.Contains(priorities, P.Name); exist == false {
		err := fmt.Errorf("priority value should be only 'high', 'normal' or 'low' got: %v", P.Name)
		return err
	}
	return nil
}

type Status struct {
	Name string
}

func (S Status) validate() bool {
	statuses := []string{"todo", "inProgress", "done"}
	if exist := slices.Contains(statuses, S.Name); exist == false {
		return false
	}
	return true
}

func TodoCreate(ctx context.Context, db *sql.DB, title string, priority Priority, status Status) (int64, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err := fmt.Errorf("failed to begin tx: %w", err)
		return 0, err
	}

	defer tx.Rollback()

	if title == "" {
		err := fmt.Errorf("title should not be empty")
		return 0, err
	}

	err = priority.validate()
	if err != nil {
		return 0, err
	}

	checkStatusValue := status.validate()
	if checkStatusValue == false {
		status.Name = "todo"
		slog.Warn("status's value is not valid, default value -> 'todo' will be applied")
	}

	res, err := tx.ExecContext(ctx, "INSERT INTO todo (title, status, priority) VALUES (?, ?, ?)", title, status.Name, priority.Name)
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

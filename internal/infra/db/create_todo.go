package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"slices"
)

func main() {

}

func handleErr(title string) {
	slog.Error(title)
	os.Exit(1)
}

type priority struct {
	name string
}

func (P priority) validate() {
	if P.name == "" {
		handleErr("priority value should be only 'high', 'normal' or 'low'")
	}
	priorities := []string{"high", "normal", "low"}
	if exist := slices.Contains(priorities, P.name); exist == false {
		handleErr("priority value should be only 'high', 'normal' or 'low'")
	}
}

type status struct {
	status string
}

func (S status) validate() bool {
	statuses := []string{"todo", "inProgress", "done"}
	if exist := slices.Contains(statuses, S.status); exist == false {
		return false
	}
	return true
}

func TodoCreate(ctx context.Context, db *sql.DB, title string, priority priority, status status) int64 {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		handleErr("failed to begin tx" + error.Error(err))
	}

	defer tx.Rollback()

	if title == "" {
		handleErr("title should not be empty")
	}
	priority.validate()
	checkStatusValue := status.validate()
	if checkStatusValue == false {
		slog.Warn("status's value is not valid, default value -> 'todo' will be applied")
	}

	res, err := tx.ExecContext(ctx, "INSERT INTO todo (title, status, priority) VALUES (?, ?, ?)", title, status, priority)
	if err != nil {
		handleErr("failed to add new todo with err: " + error.Error(err))
	}

	ID, err := res.LastInsertId()
	if err != nil {
		handleErr("failed to add new todo with err: " + error.Error(err))
	}
	return ID

}

package db

import (
	"database/sql"
	"fmt"
)

func TodoCreate(db *sql.DB, title, priority, status string) (int64, error) {

	res, err := db.Exec("INSERT INTO todo (title, status, priority) VALUES (?, ?, ?)", title, status, priority)
	if err != nil {
		err := fmt.Errorf("failed to add new todo with err: %w", err)
		return 0, err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		err := fmt.Errorf("failed to add new todo with err: %w", err)
		return 0, err
	}

	return ID, nil
}

package service

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
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

func CreateTodoService(client db.Client, title string, priority Priority, status Status) (int64, error) {
	if title == "" {
		err := fmt.Errorf("title should not be empty")
		return 0, err
	}

	err := priority.validate()
	if err != nil {
		return 0, err
	}

	checkStatusValue := status.validate()
	if checkStatusValue == false {
		status.Name = "todo"
		slog.Warn("status's value is not valid, default value -> 'todo' will be applied")
	}
	id, err := db.TodoCreate(client.DB, title, priority.Name, status.Name)
	return id, err
}

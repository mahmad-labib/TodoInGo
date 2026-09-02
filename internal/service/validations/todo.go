package service

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type todoValidate struct {
	Todo *domain.Todo
}

func (t todoValidate) validate() error {
	if t.Todo.Title == "" {
		err := fmt.Errorf("title should not be empty")
		return err
	}
	err := validatePriority(&t)
	validateStatus(&t)

	return err
}

func validatePriority(t *todoValidate) error {
	if t.Todo.Priority == "" {
		err := fmt.Errorf("priority value should be only 'high', 'normal' or 'low' got nil")
		return err
	}
	priorities := []string{"high", "normal", "low"}
	if exist := slices.Contains(priorities, t.Todo.Priority); exist == false {
		err := fmt.Errorf("priority value should be only 'high', 'normal' or 'low' got: %v", t.Todo.Priority)
		return err
	}
	return nil
}

func validateStatus(t *todoValidate) {
	statuses := []string{"todo", "inProgress", "done"}
	if exist := slices.Contains(statuses, t.Todo.Status); exist == false {
		t.Todo.Status = "todo"
		slog.Warn("status's value is not valid, default value -> 'todo' will be applied")
	}
}

func CreateTodoService(Todo *domain.Todo) error {
	TodoWithValidation := todoValidate{Todo: Todo}
	err := TodoWithValidation.validate()
	return err
}

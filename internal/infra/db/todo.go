package db

import (
	"fmt"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db/model"
	"gorm.io/gorm"
)

type Todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) *Todo {
	return &Todo{db: db}
}

func (r *Todo) Create(todo *domain.Todo) (uint, error) {
	dbTodo := model.Todo{
		Title:    todo.Title,
		Status:   todo.Status,
		Priority: todo.Priority,
	}

	result := r.db.Create(&dbTodo)
	if result.Error != nil {
		err := fmt.Errorf("failed to create new todo with err: %w", result.Error)
		return 0, err
	}

	todo.ID = dbTodo.ID
	return todo.ID, nil
}

package db

import (
	"fmt"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db/models"
	"gorm.io/gorm"
)

type TodoGormRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoGormRepo {
	return &TodoGormRepo{db: db}
}

func (r *TodoGormRepo) Create(todo *domain.Todo) (uint, error) {
	dbTodo := models.TodoModel{
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

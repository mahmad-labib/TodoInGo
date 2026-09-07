package db

import (
	"context"
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

func (r *Todo) Create(ctx context.Context, item *domain.Todo) (uint, error) {
	dbTodo := model.Todo{
		Title:    item.Title,
		Status:   item.Status,
		Priority: item.Priority,
	}

	result := r.db.Create(&dbTodo)
	if result.Error != nil {
		err := fmt.Errorf("failed to create new todo with err: %w", result.Error)
		return 0, err
	}

	return dbTodo.ID, nil
}

func (r *Todo) List(ctx context.Context, sort string) (*[]domain.TodoList, error) {
	var todosList *[]domain.TodoList

	err := r.db.
		Table("todo").
		Select("id", "title", "status", "priority", "created_at").
		Order(sort).
		Find(&todosList).
		Error

	if err != nil {
		return nil, err
	}

	return todosList, nil
}

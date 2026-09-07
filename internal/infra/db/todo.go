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

func (r *Todo) List(ctx context.Context, list *domain.List) ([]model.Todo, error) {
	var todosList []model.Todo
	err := r.db.Select("id", "title", "status", "priority", "created_at").Order(list.Sort).Find(&todosList).Error
	if err != nil {
		return nil, err
	}
	return todosList, nil
}

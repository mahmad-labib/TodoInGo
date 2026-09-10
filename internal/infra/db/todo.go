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

	result := r.db.WithContext(ctx).Create(&dbTodo)
	if result.Error != nil {
		err := fmt.Errorf("failed to create new todo with err: %w", result.Error)
		return 0, err
	}

	return dbTodo.ID, nil
}

func (r *Todo) List(ctx context.Context, sort string) ([]domain.TodoList, error) {
	var todos []model.Todo

	err := r.db.
		WithContext(ctx).
		Order(sort).
		Find(&todos).
		Error

	if err != nil {
		return nil, err
	}

	list := make([]domain.TodoList, 0, len(todos))

	for _, t := range todos {
		list = append(list, domain.TodoList{
			ID:        t.ID,
			Title:     t.Title,
			Status:    t.Status,
			Priority:  t.Priority,
			CreatedAt: t.CreatedAt,
		})
	}

	return list, nil
}

func (r *Todo) Update(ctx context.Context, item *domain.Todo) (*domain.Todo, error) {

	var todo model.Todo
	result := r.db.WithContext(ctx).First(&todo, item.ID)
	if result.Error != nil {
		return item, result.Error
	}

	todo.Status = item.Status

	result = r.db.WithContext(ctx).Save(&todo)
	if result.Error != nil {
		return item, result.Error
	}

	item.ID = todo.ID
	item.Title = todo.Title
	item.Status = todo.Status
	item.Priority = todo.Priority

	return item, nil
}

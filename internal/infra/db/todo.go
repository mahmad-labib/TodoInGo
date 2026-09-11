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

func (r *Todo) UpdateStatus(ctx context.Context, item *domain.Todo) error {
	res := r.db.WithContext(ctx).Model(&model.Todo{}).Where("id = ?", item.ID).Update("status", item.Status)

	if res.Error != nil {
		return fmt.Errorf("error when update todo id=%d: %w", item.ID, res.Error)
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("todo: id=%d: is not exist", item.ID)
	}

	return nil
}

func (r *Todo) FindByID(ctx context.Context, id uint) (*domain.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}

	var item = domain.Todo{
		ID:       todo.ID,
		Title:    todo.Title,
		Status:   todo.Status,
		Priority: todo.Priority,
	}

	return &item, nil
}

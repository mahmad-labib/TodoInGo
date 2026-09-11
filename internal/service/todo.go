package service

import (
	"context"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	UpdateStatus(ctx context.Context, todo *domain.Todo) error
	FindByID(ctx context.Context, id uint) (*domain.Todo, error)
}

type TodoValidator interface {
	Validate(item *domain.Todo) error
	ValidateStatus(status string) error
	ValidateSort(sort *string) error
}

type Todo struct {
	repo      TodoRepository
	validator TodoValidator
}

func NewTodo(repo TodoRepository, validator TodoValidator) *Todo {
	return &Todo{repo: repo, validator: validator}
}

func (s *Todo) Create(ctx context.Context, item *domain.Todo) (uint, error) {
	if err := s.validator.Validate(item); err != nil {
		return 0, err
	}

	return s.repo.Create(ctx, item)
}

func (s *Todo) List(ctx context.Context, sort string) ([]domain.TodoList, error) {
	err := s.validator.ValidateSort(&sort)
	if err != nil {
		return nil, err
	}

	sortedList, err := s.repo.List(ctx, sort)
	if err != nil {
		return sortedList, err
	}

	return sortedList, nil
}

func (s *Todo) UpdateStatus(ctx context.Context, item *domain.Todo) (*domain.Todo, error) {
	if err := s.validator.ValidateStatus(item.Status); err != nil {
		return item, err
	}

	err := s.repo.UpdateStatus(ctx, item)
	if err != nil {
		return item, err
	}

	updatedItem, err := s.repo.FindByID(ctx, item.ID)
	if err != nil {
		return updatedItem, err
	}

	return updatedItem, nil
}

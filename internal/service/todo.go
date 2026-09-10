package service

import (
	"context"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)
}

type TodoValidator interface {
	Validate(item *domain.Todo) error
	ValidateStatus(item *domain.Todo) error
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
	sortedList, err := s.repo.List(ctx, sort)
	if err != nil {
		return sortedList, err
	}

	return sortedList, nil
}

func (s *Todo) Update(ctx context.Context, item *domain.Todo) (*domain.Todo, error) {
	if err := s.validator.ValidateStatus(item); err != nil {
		return item, err
	}

	item, err := s.repo.Update(ctx, item)
	if err != nil {
		return item, err
	}

	return item, nil
}

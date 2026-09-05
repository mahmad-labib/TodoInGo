package service

import (
	"context"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) (uint, error)
}

type TodoValidator interface {
	Validate(item *domain.Todo) error
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

package service

import (
	"context"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) (uint, error)
	List(ctx context.Context, sort *domain.Sort) ([]model.Todo, error)
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

func (s *Todo) List(ctx context.Context, sort *domain.Sort) (domain.List, error) {
	sortedList, err := s.repo.List(ctx, sort)
	List := domain.List{
		Data: sortedList,
	}
	return List, err
}

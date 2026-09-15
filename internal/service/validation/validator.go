package validation

import (
	"errors"
	"slices"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type Validator struct{}

var (
	ErrInvalidTitle    = errors.New("title is required")
	ErrTitleLength     = errors.New("title cannot be longer than 50 characters")
	ErrInvalidPriority = errors.New("priority cant be empty, value should be high, normal or low")
	ErrPriorityLength  = errors.New(("priority cant be empty, value should be high, normal or low"))
	ErrInvalidStatus   = errors.New("status value should be todo, inProgress or done")
	ErrStatusLength    = errors.New("status cannot be empty")
	ErrInvalidSort     = errors.New("sort should be status, created_at, priority")
)

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(item *domain.Todo) error {
	err1 := v.validateTitle(item.Title)

	err2 := v.ValidateStatus(item.Status)

	err3 := v.validatePriority(item.Priority)

	return errors.Join(err1, err2, err3)
}

func (v *Validator) validateTitle(title string) error {
	if title == "" {
		return ErrInvalidTitle
	}

	if len(title) > 50 {
		return ErrTitleLength
	}

	return nil
}

func (v *Validator) ValidateStatus(status string) error {
	statuses := []string{"todo", "inProgress", "done"}

	if status == "" {
		return ErrStatusLength
	}

	if !slices.Contains(statuses, status) {
		return ErrInvalidStatus
	}

	return nil
}
func (v *Validator) validatePriority(priority string) error {
	priorities := []string{"high", "normal", "low"}

	if priority == "" {
		return ErrPriorityLength
	}

	if slices.Contains(priorities, priority) {
		return nil
	}

	return ErrInvalidPriority
}

func (v *Validator) ValidateSort(sort *string) error {
	sortList := []string{"status", "created_at", "priority"}

	if *sort == "" {
		*sort = "created_at"
	}

	if slices.Contains(sortList, *sort) {
		return nil
	}

	return ErrInvalidSort

}

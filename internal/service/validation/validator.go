package validation

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(item *domain.Todo) error {
	err1 := v.validateTitle(item.Title)

	err2 := v.validateStatus(item.Status)

	err3 := v.validatePriority(item.Priority)

	return errors.Join(err1, err2, err3)
}

func (v *Validator) ValidateStatus(item *domain.Todo) error {
	err := v.validateStatus(item.Status)

	return err
}

func (v *Validator) validateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if len(title) > 50 {
		return fmt.Errorf("title cannot be longer than 50 characters")
	}

	return nil
}

func (v *Validator) validateStatus(status string) error {
	statuses := []string{"todo", "inProgress", "done"}

	if status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	if !slices.Contains(statuses, status) {
		return fmt.Errorf("status value should be todo, inProgress or done")
	}

	return nil
}
func (v *Validator) validatePriority(priority string) error {
	priorities := []string{"high", "normal", "low"}

	if priority == "" {
		return fmt.Errorf("priority cant be empty, value should be high, normal or low")
	}

	if slices.Contains(priorities, priority) {
		return nil
	}

	return fmt.Errorf("priority should be high, normal or low")
}

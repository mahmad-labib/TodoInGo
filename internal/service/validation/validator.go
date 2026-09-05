package validation

import (
	"errors"
	"fmt"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(item *domain.Todo) error {
	err1 := v.validateTitle(item.Title)

	err2 := v.validateStatus(item.Status)

	return errors.Join(err1, err2)
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
	if status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	return nil
}

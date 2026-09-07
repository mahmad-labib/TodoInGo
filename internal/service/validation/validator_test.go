package validation

import (
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
)

func TestServiceTodo_Success(t *testing.T) {
	Item := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	err := New().Validate(&Item)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestServiceTodo_EmptyTitle(t *testing.T) {
	Item := domain.Todo{
		Title:    "",
		Priority: "high",
		Status:   "todo",
	}

	err := New().Validate(&Item)
	if err == nil {
		t.Fatalf("expected error, got no error")
	}
}

func TestServiceTodo_EmptyPriority(t *testing.T) {
	Item := domain.Todo{
		Title:    "Test",
		Priority: "",
		Status:   "todo",
	}

	err := New().Validate(&Item)
	if err == nil {
		t.Fatalf("expected error, got no error")
	}
}

func TestServiceTodo_InvalidPriority(t *testing.T) {
	Item := domain.Todo{
		Title:    "Test",
		Priority: "Invalid",
		Status:   "todo",
	}

	err := New().Validate(&Item)
	if err == nil {
		t.Fatalf("expected error, got no error")
	}
}

func TestServiceTodo_EmptyStatus(t *testing.T) {
	Item := domain.Todo{
		Title:    "Test",
		Priority: "low",
		Status:   "",
	}

	err := New().Validate(&Item)
	if err == nil {
		t.Fatalf("expected error, got no error")
	}
}

func TestServiceTodo_InavalidStatus(t *testing.T) {
	Item := domain.Todo{
		Title:    "Test",
		Priority: "low",
		Status:   "Invalid",
	}

	err := New().Validate(&Item)
	if err == nil {
		t.Fatalf("expected error, got no error")
	}
}

package service_test

import (
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	service "github.com/mahmad-labib/TodoInGo/internal/service/validations"
	"gorm.io/gorm"
)

func returnClient() (*gorm.DB, error) {
	client, err := db.NewClient(&config.DB{Name: "test"})
	sqliteDB, err := client.DB()
	if err != nil {
		return client, err
	}
	_, err = db.MigrateUp(sqliteDB)
	return client, err
}

func TestAddTodo_success(t *testing.T) {
	client, err := returnClient()
	TodoRepo := db.NewTodo(client)

	Todo := domain.Todo{
		Title:    "Test",
		Status:   "todo",
		Priority: "high",
	}
	validatedTodo, err := service.CreateTodoService(&Todo)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	id, err := TodoRepo.Create(validatedTodo.Todo)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected returned ID, got nil")
	}
}

func TestAddTodo_EmptyTitle(t *testing.T) {
	Todo := domain.Todo{
		Title:    "",
		Status:   "todo",
		Priority: "high",
	}

	_, err := service.CreateTodoService(&Todo)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddTodo_UniqueTitle(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	TodoRepo := db.NewTodo(client)

	Todo := domain.Todo{
		Title:    "Test",
		Status:   "todo",
		Priority: "high",
	}

	validatedTodo, err := service.CreateTodoService(&Todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	id, err := TodoRepo.Create(validatedTodo.Todo)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected returned ID, got nil")
	}

	id, err = TodoRepo.Create(validatedTodo.Todo)
	if err == nil {
		t.Fatal("expected Error, got no error")
	}
	if id > 0 {
		t.Fatalf("expected no ID, got: %v", id)
	}
}

func TestAddTodo_EmptyPriority(t *testing.T) {

	Todo := domain.Todo{
		Title:    "Test",
		Status:   "todo",
		Priority: "",
	}

	_, err := service.CreateTodoService(&Todo)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddTodo_EmptyStatus(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	TodoRepo := db.NewTodo(client)

	Todo := domain.Todo{
		Title:    "Test",
		Status:   "",
		Priority: "high",
	}

	validatedTodo, err := service.CreateTodoService(&Todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	id, err := TodoRepo.Create(validatedTodo.Todo)

	if err != nil {
		t.Fatalf("expected no Error, got: %v", err)
	}
	if id == 0 {
		t.Fatalf("expected ID, got nil")
	}
}

func TestAddTodo_InvalidPriority(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Status:   "todo",
		Priority: "hi",
	}

	_, err := service.CreateTodoService(&Todo)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddTodo_InvalidStatus(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	TodoRepo := db.NewTodo(client)

	Todo := domain.Todo{
		Title:    "Test",
		Status:   "to",
		Priority: "high",
	}

	validatedTodo, err := service.CreateTodoService(&Todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	id, err := TodoRepo.Create(validatedTodo.Todo)

	if err != nil {
		t.Fatalf("expected no Error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected ID, got nil")
	}
}

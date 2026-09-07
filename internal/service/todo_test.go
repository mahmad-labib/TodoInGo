package service

import (
	"context"
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
)

func TestDBClient(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error gor %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	db.MigrateUp(DB)

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}
}

func TestDBClient_TestValidation(t *testing.T) {
	Todo := domain.Todo{
		Title:    "",
		Priority: "",
		Status:   "",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error gor %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	db.MigrateUp(DB)

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err == nil {
		t.Fatalf("expected an error got no error")
	}

	if id > 0 {
		t.Fatalf("expected no id got: %v", id)
	}
}

func TestGetTodoList_Populated(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error gor %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	db.MigrateUp(DB)

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	sort := "title"
	List, err := service.List(ctx, sort)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}
	if len(*List) == 0 {
		t.Fatal("expected a list got nil")
	}

}

func TestGetTodoList_Empty(t *testing.T) {
	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error gor %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}

	db.MigrateUp(DB)

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	sort := "title"
	List, err := service.List(ctx, sort)

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}
	if len(*List) > 0 {
		t.Fatalf("expected nil got: %v", List)
	}

}

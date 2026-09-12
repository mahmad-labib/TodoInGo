package service

import (
	"context"
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
)

func TestTodo_Create(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}
}

func TestTodo_Create_InvalidData(t *testing.T) {
	Todo := domain.Todo{
		Title:    "",
		Priority: "",
		Status:   "",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

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

func TestTodo_List_Populated(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	sort := "status"
	List, err := service.List(ctx, sort)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}
	if len(List) == 0 {
		t.Fatal("expected a list got nil")
	}

}

func TestTodo_List_Empty(t *testing.T) {
	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	sort := "status"
	List, err := service.List(ctx, sort)

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}
	if len(List) > 0 {
		t.Fatalf("expected nil got: %v", List)
	}

}

func TestTodo_Update(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	item := &domain.Todo{
		ID:     id,
		Status: "done",
	}

	todo, err := service.UpdateStatus(ctx, item)
	if err != nil {
		t.Fatalf("expected no error go %v", err)
	}
	if todo.ID != 1 {
		t.Fatalf("expected todo with id 1 got %v", todo.ID)
	}
	if todo.Status != "done" {
		t.Fatalf("expected status to be done got %v", todo.Status)
	}
}

func TestTodo_Update_InvalidID(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	item := &domain.Todo{
		ID:     0,
		Status: "done",
	}

	_, err = service.UpdateStatus(ctx, item)
	if err == nil {
		t.Fatalf("expected error got nil")
	}
}

func TestTodo_Delete(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	err = service.Delete(ctx, id)
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = service.repo.FindByID(ctx, id)
	if err == nil {
		t.Fatalf("expected error got nil")
	}

}

func TestTodo_Delete_InvalidID(t *testing.T) {
	Todo := domain.Todo{
		Title:    "Test",
		Priority: "high",
		Status:   "todo",
	}

	dbClient, err := db.NewClient(&config.DB{Name: "test"})
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	DB, err := dbClient.DB()
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	_, err = db.MigrateUp(DB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	validator := validation.New()
	repo := db.NewTodo(dbClient)
	service := NewTodo(repo, validator)

	ctx := context.Background()

	id, err := service.Create(ctx, &Todo)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

	if id == 0 {
		t.Fatal("expected id got no id")
	}

	err = service.Delete(ctx, 0)
	if err == nil {
		t.Fatalf("expected error got nil")
	}

	_, err = service.repo.FindByID(ctx, id)
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}

}

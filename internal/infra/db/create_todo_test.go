package db_test

import (
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func returnClient() (db.Client, error) {
	client, err := db.NewClient(":memory:")
	_, err = db.MigrateUp(client.DB)
	return *client, err
}

func TestAddTodo_success(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "todo"}
	priority := db.Priority{Name: "high"}
	id, err := db.TodoCreate(ctx, client.DB, "Todo", priority, status)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected returned ID, got nil")
	}
}

func TestAddTodo_EmptyTitle(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "todo"}
	priority := db.Priority{Name: "high"}
	id, err := db.TodoCreate(ctx, client.DB, "", priority, status)
	if err == nil {
		t.Fatal("expected Error, got no error")
	}
	if id > 0 {
		t.Fatalf("expected no ID, got: %v", id)
	}
}

func TestAddTodo_UniqueTitle(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "todo"}
	priority := db.Priority{Name: "high"}
	id, err := db.TodoCreate(ctx, client.DB, "title", priority, status)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected returned ID, got nil")
	}
	id, err = db.TodoCreate(ctx, client.DB, "title", priority, status)
	if err == nil {
		t.Fatal("expected Error, got no error")
	}
	if id > 0 {
		t.Fatalf("expected no ID, got: %v", id)
	}
}

func TestAddTodo_EmptyPriority(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "todo"}
	priority := db.Priority{}
	id, err := db.TodoCreate(ctx, client.DB, "title", priority, status)
	if err == nil {
		t.Fatal("expected Error, got no error")
	}
	if id > 0 {
		t.Fatalf("expected no ID, got: %v", id)
	}
}

func TestAddTodo_EmptyStatus(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: ""}
	priority := db.Priority{Name: "high"}
	id, err := db.TodoCreate(ctx, client.DB, "title", priority, status)

	if err != nil {
		t.Fatalf("expected no Error, got: %v", err)
	}
	if id == 0 {
		t.Fatalf("expected ID, got nil")
	}
}

func TestAddTodo_InvalidPriority(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "todo"}
	priority := db.Priority{Name: "severe"}
	id, err := db.TodoCreate(ctx, client.DB, "title", priority, status)

	if err == nil {
		t.Fatal("expected Error, got no Error")
	}
	if id > 0 {
		t.Fatalf("expected No ID, got %v", id)
	}
}

func TestAddTodo_InvalidStatus(t *testing.T) {
	client, err := returnClient()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()

	ctx := t.Context()
	status := db.Status{Name: "will do"}
	priority := db.Priority{Name: "high"}
	id, err := db.TodoCreate(ctx, client.DB, "title", priority, status)

	if err != nil {
		t.Fatalf("expected no Error, got: %v", err)
	}
	if id == 0 {
		t.Fatal("expected ID, got nil")
	}
}

package db_test

import (
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

// TestNewClient_Success tests valid connection creation
func TestNewClient_Success(t *testing.T) {
	// Use in-memory SQLite for tests
	client, err := db.NewClient(&config.DB{Name: "test"})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if client == nil {
		t.Fatal("expected DB instance to not be nil")
	}
}

func TestNewClient_EmptyDSN(t *testing.T) {
	client, err := db.NewClient(&config.DB{})

	if err == nil {
		t.Fatal("expected error for empty dsn, got nil")
	}
	if client != nil {
		t.Fatalf("expected nil client, got: %v", client)
	}
}

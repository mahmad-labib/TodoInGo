package db

import (
	"fmt"
	"testing"
)

// TestNewClient_Success tests valid connection creation
func TestNewClient_Success(t *testing.T) {
	// Use in-memory SQLite for tests
	client, msg, err := NewClient(":memory:")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer client.DB.Close()
	fmt.Println(msg)
	if client.DB == nil {
		t.Fatal("expected DB instance to not be nil")
	}
}

func TestNewClient_EmptyDSN(t *testing.T) {
	client, _, err := NewClient("")
	if err == nil {
		t.Fatal("expected error for empty dsn, got nil")
	}
	if client != nil {
		t.Fatalf("expected nil client, got: %v", client)
	}
}

func Test_Migration(t *testing.T) {
	client, _, err := NewClient(":memory:")
	if err != nil {
		t.Fatalf("migration setup expected no error, got: %v", err)
	}
	defer client.DB.Close()

	msg, err := MigrateUp(client.DB)

	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	if msg == "" {
		t.Fatal("expected msg got empty string")
	}

	var tableName string
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='todo';"
	err = client.DB.QueryRow(query).Scan(&tableName)

	if err != nil {
		t.Fatalf("expected todo table go: %v", err)
	}

	if tableName != "todo" {
		t.Fatalf("expected todo table got: %s", tableName)
	}

	msg, err = MigrateDown(client.DB)

	if err != nil {
		t.Fatalf("migration Down failed: %v", err)
	}

	if msg == "" {
		t.Fatal("expected msg got empty string")
	}

}

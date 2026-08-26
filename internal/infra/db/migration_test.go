package db_test

import (
	"testing"

	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func Test_Migration(t *testing.T) {
	client, err := db.NewClient(":memory:")
	if err != nil {
		t.Fatalf("migration setup expected no error, got: %v", err)
	}
	defer client.DB.Close()

	msg, err := db.MigrateUp(client.DB)

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

	msg, err = db.MigrateDown(client.DB)

	if err != nil {
		t.Fatalf("migration Down failed: %v", err)
	}

	if msg == "" {
		t.Fatal("expected msg got empty string")
	}

}

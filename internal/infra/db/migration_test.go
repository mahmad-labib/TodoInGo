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
	sqliteDB, err := client.DB.DB()
	if err != nil {
		t.Fatalf("fetching sqliteDB expected no error, got: %v", err)
	}

	msg, err := db.MigrateUp(sqliteDB)
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}
	if msg == "" {
		t.Fatal("expected msg got empty string")
	}

	var tableName string
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='todo';"
	err = sqliteDB.QueryRow(query).Scan(&tableName)

	if err != nil {
		t.Fatalf("expected todo table go: %v", err)
	}

	if tableName != "todo" {
		t.Fatalf("expected todo table got: %s", tableName)
	}

	msg, err = db.MigrateDown(sqliteDB)

	if err != nil {
		t.Fatalf("migration Down failed: %v", err)
	}

	if msg == "" {
		t.Fatal("expected msg got empty string")
	}

}

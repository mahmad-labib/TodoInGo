package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func main() {
	defaultPath := getDBFile()

	client, err := db.NewClient(defaultPath)
	handleErr(err)
	defer client.DB.Close()
	handleResult("connection established successfully")

	rollback := checkMigrationStatus()
	if rollback {
		result, err := db.MigrateDown(client.DB)
		handleErr(err)
		handleResult(result)
	} else {
		result, err := db.MigrateUp(client.DB)
		handleErr(err)
		handleResult(result)
	}
}

func handleErr(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func handleResult(result string) {
	if result != "" {
		slog.Info(result)
	}
}

func checkMigrationStatus() bool {
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if v == "rollback" {
				return true
			}
		}
	}
	return false
}

func getDBFile() string {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, will use default values")
	}

	dbName := os.Getenv("DB_NAME")
	dbPath := os.Getenv("DB_PATH")

	if dbName == "" {
		dbName = "myApp.db"
	}
	if dbPath == "" {
		dbPath = "."
	}

	defaultPath := filepath.Join(dbPath, dbName)
	return defaultPath
}

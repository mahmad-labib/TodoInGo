package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetDBFile() string {
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

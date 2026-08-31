package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func GetDBFile() *db.DB {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, will use default values")
	}

	cfg := db.DB{
		DBName: os.Getenv("DB_NAME"),
		DBPath: os.Getenv("DB_PATH"),
	}

	return &cfg
}

// config/config.go

package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DB
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, will use default values")
	}

	return &Config{
		DB: NewDB(),
	}, nil
}

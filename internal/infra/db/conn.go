package db

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/mahmad-labib/TodoInGo/internal/config"
	"gorm.io/gorm"
)

func NewClient(cfg *config.DB) (*gorm.DB, error) {
	dns := resolveFilePath(cfg)

	if dns == "" {
		return nil, errors.New("DSN cannot be empty")
	}

	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database : %w", err)
	}

	return db, nil
}

func resolveFilePath(cfg *config.DB) string {
	if cfg.Name == "test" {
		return ":memory:"
	}
	return filepath.Join(cfg.Path, cfg.Name)
}

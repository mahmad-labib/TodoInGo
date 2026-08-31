package db

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/mahmad-labib/TodoInGo/internal/config"
	"gorm.io/gorm"
)

// Client wraps sql.DB
type Client struct {
	DB *gorm.DB
}

func NewClient(dsn string) (*Client, error) {
	if dsn == "" {
		return nil, errors.New("DSN cannot be empty")
	}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database : %w", err)
	}

	return &Client{DB: db}, nil
}

func ResolveFilePath(cfg *config.DB) string {
	name := cfg.Name
	path := cfg.Path

	if name == "" {
		name = "myApp.db"
	}
	if path == "" {
		path = "."
	}

	return filepath.Join(cfg.Path, cfg.Name)
}

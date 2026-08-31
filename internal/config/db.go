package config

import (
	"os"
)

type DB struct {
	Name string
	Path string
}

func NewDB() DB {
	return DB{
		Name: os.Getenv("DB_NAME"),
		Path: os.Getenv("DB_PATH"),
	}
}

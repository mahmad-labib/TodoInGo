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
		Name: getDBName(),
		Path: getDBPath(),
	}
}

func getDBName() string {
	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		dbName = "myApp.db"
	}

	return dbName
}

func getDBPath() string {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		dbPath = "."
	}

	return dbPath
}

package main

import (
	"log/slog"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Failed to load config", "err: ", err)
	}
	defaultPath := db.ResolveFilePath(&cfg.DB)
	client, err := db.NewClient(defaultPath)
	handleErr(err)
	handleResult("connection established successfully")

	rollback := checkMigrationStatus()
	if rollback {
		sqliteDB, err := client.DB.DB()
		handleErr(err)
		result, err := db.MigrateDown(sqliteDB)
		handleErr(err)
		handleResult(result)
	} else {
		sqliteDB, err := client.DB.DB()
		handleErr(err)
		result, err := db.MigrateUp(sqliteDB)
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

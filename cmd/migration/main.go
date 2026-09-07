package main

import (
	"log/slog"
	"os"
	"slices"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Failed to load config", "err: ", err)

		os.Exit(1)
	}

	client, err := db.NewClient(&cfg.DB)
	if err != nil {
		slog.Error(err.Error())

		os.Exit(1)
	}

	status := checkMigrationStatus()
	if status == "rollback" {
		err := migrationDown(client)
		if err != nil {
			slog.Error(err.Error())

			os.Exit(1)
		}

		slog.Info("Rollback successful")

		return
	}

	err = migrationUp(client)
	if err != nil {
		slog.Error(err.Error())

		os.Exit(1)
	}

	slog.Info("Migration successful")
}

func migrationDown(client *gorm.DB) error {
	sqliteDB, err := client.DB()
	if err != nil {
		return err
	}

	_, err = db.MigrateDown(sqliteDB)
	if err != nil {
		return err
	}

	return nil
}

func migrationUp(client *gorm.DB) error {
	sqliteDB, err := client.DB()
	if err != nil {
		return err
	}

	_, err = db.MigrateUp(sqliteDB)
	if err != nil {
		return err
	}

	return nil
}

func checkMigrationStatus() string {
	if len(os.Args) > 1 {
		if slices.Contains(os.Args[1:], "rollback") {
			return "rollback"
		}
	}

	return "migrate"
}

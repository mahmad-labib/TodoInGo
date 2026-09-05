package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/cli"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	"github.com/mahmad-labib/TodoInGo/internal/service"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to load config %v", err))

		os.Exit(1)
	}

	dbClient, err := db.NewClient(&cfg.DB)
	if err != nil {
		slog.Error(fmt.Sprintf("couldn't establish connection %v", err))

		os.Exit(1)
	}

	repo := db.NewTodo(dbClient)

	validator := validation.New()

	service := service.NewTodo(repo, validator)

	ctx := context.Background()

	if err := cli.NewTodo(service).Run(ctx); err != nil {
		slog.Error(fmt.Sprintf("Error executing command %v", err))

		os.Exit(1)
	}

	slog.Info("Command executed successfully")
}

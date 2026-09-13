package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mahmad-labib/TodoInGo/http/router"
	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	"github.com/mahmad-labib/TodoInGo/internal/service"
	"github.com/mahmad-labib/TodoInGo/internal/service/validation"
)

type Router struct{}

func main() {

	service := initiateService()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.Mount("/v1", router.NewTodoRouter(service))

	done := make(chan struct{})

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		fmt.Println("Shutdown signal received, finishing in-flight requests...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Shutdown error: %v\n", err)
		}

		close(done)
	}()

	fmt.Println("Server running on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Server error: %v\n", err)
	}

	<-done
	fmt.Println("Server exited cleanly")
}

func initiateService() *service.Todo {
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
	return service
}

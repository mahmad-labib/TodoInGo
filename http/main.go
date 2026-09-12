package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

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

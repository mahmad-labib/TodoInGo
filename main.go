package main

import (
	"log/slog"
	"os"

	"github.com/mahmad-labib/TodoInGo/cmd/todo"
)

func main() {
	id, err := todo.CmdTodo()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Todo created with", "id", id)
}

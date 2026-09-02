package cli

import (
	"flag"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
	service "github.com/mahmad-labib/TodoInGo/internal/service/validations"
)

func HandleCmd(client *db.Client) (uint, error) {
	TodoRepo := db.NewTodoRepo(client.DB)

	Todo := GetCmdFlags()
	validatedTodo, err := service.CreateTodoService(Todo)
	if err != nil {
		return 0, err
	}

	id, err := TodoRepo.Create(validatedTodo.Todo)
	return id, err
}

func GetCmdFlags() *domain.Todo {
	Todo := domain.Todo{}

	flag.StringVar(&Todo.Title, "title", "", "todo title")
	flag.StringVar(&Todo.Priority, "priority", "", "todo priority ['high', 'normal', 'low']")
	flag.StringVar(&Todo.Status, "status", "", "todo status ['todo', 'inProgress', 'done']")

	flag.Parse()

	return &Todo
}

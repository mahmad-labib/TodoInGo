package cli

import (
	"context"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	validation "github.com/urfave/cli-validation"
	"github.com/urfave/cli/v3"
)

type TodoService interface {
	Create(ctx context.Context, item *domain.Todo) (uint, error)
	List(ctx context.Context, sort string) ([]domain.TodoList, error)
	Update(ctx context.Context, item *domain.Todo) (*domain.Todo, error)
}

type Todo struct {
	service TodoService
}

func NewTodo(service TodoService) *Todo {
	return &Todo{service: service}
}

func (c *Todo) Run(ctx context.Context) error {
	cmd := &cli.Command{
		Usage: "Todo application CLI",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create a new todo item",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Aliases:  []string{"t"},
						Required: true,
						Usage:    "Title of the new todo item",
					},
					&cli.StringFlag{
						Name:     "status",
						Aliases:  []string{"s"},
						Required: true,
						Usage:    "Status of the new todo item",
					},
					&cli.StringFlag{
						Name:     "priority",
						Aliases:  []string{"p"},
						Required: true,
						Usage:    "Priority of the new todo item",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return c.createAction(ctx, cmd)
				},
			},
			{
				Name:  "list",
				Usage: "get a list of todo items",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "sort",
						Required:  false,
						Usage:     "sort todo list by status,priority,created_at",
						Validator: validation.Enum("status", "priority", "created_at"),
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return c.listAction(ctx, cmd)
				},
			},
			{
				Name:  "update",
				Usage: "update an item status",
				Flags: []cli.Flag{
					&cli.UintFlag{
						Name:     "id",
						Required: true,
						Usage:    "id of the item to update",
					},
					&cli.StringFlag{
						Name:      "status",
						Required:  true,
						Usage:     "the new item status",
						Validator: validation.Enum("todo", "inProgress", "done"),
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return c.updateAction(ctx, cmd)
				},
			},
		},
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		return err
	}

	return nil
}

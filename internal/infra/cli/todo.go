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
	List(ctx context.Context, sort *domain.Sort) (domain.List, error)
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
					c.listAction(ctx, cmd)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		return err
	}

	return nil
}

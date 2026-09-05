package cli

import (
	"context"
	"fmt"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/urfave/cli/v3"
)

func (c *Todo) createAction(ctx context.Context, cmd *cli.Command) error {
	item := &domain.Todo{
		Title:    cmd.String("title"),
		Status:   cmd.String("status"),
		Priority: cmd.String("priority"),
	}

	ID, err := c.service.Create(ctx, item)
	if err != nil {
		return err
	}

	fmt.Println("Todo item created with ID:", ID)

	return nil
}

package cli

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/urfave/cli/v3"
)

func (c *Todo) listAction(ctx context.Context, cmd *cli.Command) {
	list := &domain.List{
		Sort: cmd.String("sort"),
	}
	todosList, err := c.service.List(ctx, list)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if len(todosList) == 0 {
		slog.Info("no records found")
	}
	for _, todo := range todosList {
		fmt.Printf("%+v\n", todo)
	}
}

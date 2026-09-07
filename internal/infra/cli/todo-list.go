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
	sort := &domain.Sort{
		Sort: cmd.String("sort"),
	}
	todosList, err := c.service.List(ctx, sort)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if len(todosList.Data) == 0 {
		slog.Info("no records found")
	}
	for _, todo := range todosList.Data {
		fmt.Printf("%+v\n", todo)
	}
}

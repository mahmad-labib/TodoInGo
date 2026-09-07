package cli

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/urfave/cli/v3"
)

func (c *Todo) listAction(ctx context.Context, cmd *cli.Command) error {
	todosList, err := c.service.List(ctx, cmd.String("sort"))
	if err != nil {
		return err
	}

	if len(*todosList) == 0 {
		slog.Info("no records found")
	}

	for _, todo := range *todosList {
		fmt.Printf("%+v\n", todo)
	}

	return nil
}

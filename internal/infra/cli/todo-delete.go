package cli

import (
	"context"
	"log/slog"

	"github.com/urfave/cli/v3"
)

func (c *Todo) deleteAction(ctx context.Context, cmd *cli.Command) error {
	if err := c.service.Delete(ctx, cmd.Uint("id")); err != nil {
		return err
	}

	slog.Info("item has been deleted")

	return nil
}

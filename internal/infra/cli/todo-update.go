package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mahmad-labib/TodoInGo/internal/domain"
	"github.com/urfave/cli/v3"
)

func (c *Todo) updateAction(ctx context.Context, cmd *cli.Command) error {
	item := &domain.Todo{
		ID:     cmd.Uint("id"),
		Status: cmd.String("status"),
	}

	item, err := c.service.UpdateStatus(ctx, item)
	if err != nil {
		return err
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tTITLE\tSTATUS\tPRIORITY")
	fmt.Fprintln(writer, "--\t-----\t------\t--------")
	fmt.Fprintf(writer, "%d\t%s\t%s\t%s\n",
		item.ID,
		item.Title,
		item.Status,
		item.Priority,
	)
	writer.Flush()

	return nil
}

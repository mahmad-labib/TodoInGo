package cli

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli/v3"
)

func (c *Todo) listAction(ctx context.Context, cmd *cli.Command) error {
	todosList, err := c.service.List(ctx, cmd.String("sort"))
	if err != nil {
		return err
	}

	if len(todosList) == 0 {
		slog.Info("no records found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tTITLE\tSTATUS\tPRIORITY\tCREATED AT")
	fmt.Fprintln(writer, "--\t-----\t------\t--------\t----------")
	for _, todo := range todosList {
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n",
			todo.ID,
			todo.Title,
			todo.Status,
			todo.Priority,
			todo.CreatedAt.Format("2006-01-02 15:04"),
		)
	}
	writer.Flush()

	return nil
}

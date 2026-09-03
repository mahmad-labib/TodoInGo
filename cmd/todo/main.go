package todo

import (
	"fmt"
	"os"
	"strings"

	"github.com/mahmad-labib/TodoInGo/internal/config"
	"github.com/mahmad-labib/TodoInGo/internal/infra/cli"
	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func CmdTodo() (uint, error) {
	cfg, err := config.New()
	if err != nil {
		err = fmt.Errorf("Failed to load config %v", err)
		return 0, err
	}
	client, err := db.NewClient(&cfg.DB)
	if err != nil {
		err = fmt.Errorf("couldn't establish connection %v", err)
		return 0, err
	}

	if len(os.Args) > 1 {
		id, err := cli.HandleCmd(client)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				err = fmt.Errorf("title should be unique a title with the same name had been found")
				return 0, err
			}
			return 0, err
		}
		if id > 0 {
			return id, nil
		}
	}
	return 0, nil
}

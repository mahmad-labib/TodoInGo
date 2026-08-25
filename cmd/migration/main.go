package main

import (
	"fmt"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func main() {
	client, err := db.NewClient("./myApp.db")
	if err != nil {
		handleErr(err)
	}
	defer client.DB.Close()
	fmt.Println("connection established successfully")
	up, down := migrationFlag()
	if up {
		result, err := db.MigrateUp(client.DB)
		handleErr(err)
		if result != "" {
			fmt.Println(result)
		}
	}
	if down {
		result, err := db.MigrateDown(client.DB)
		handleErr(err)
		if result != "" {
			fmt.Println(result)
		}
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func migrationFlag() (bool, bool) {
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if v == "rollback" {
				return false, true
			}
		}
	}
	return true, false
}

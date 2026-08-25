package main

import (
	"fmt"
	"os"

	"github.com/mahmad-labib/TodoInGo/internal/infra/db"
)

func main() {
	client, err := db.NewClient("./myApp.db")
	if err != nil {
		if client != nil && client.DB != nil {
			client.DB.Close()
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer client.DB.Close()
	fmt.Println("connection established successfully")
}

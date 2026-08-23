package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mahmad-labib/TodoInGo/db"
)

func main() {
	up, down := migrationFlag()
	client, msg, err := db.NewClient("./myApp.db")
	if err != nil {
		if client.DB != nil {
			client.DB.Close()
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer client.DB.Close()
	if up {
		result, err := db.MigrateUp(client.DB)
		handleErr(err)
		msg = append(msg, result)
	}
	if down {
		result, err := db.MigrateDown(client.DB)
		handleErr(err)
		msg = append(msg, result)
	}
	for _, v := range msg {
		if v == "" {
			continue
		}
		fmt.Println(v)
	}

}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func migrationFlag() (bool, bool) {
	if len(os.Args[:]) > 2 {
		migrationFlag := flag.NewFlagSet("migration", flag.ExitOnError)
		up := migrationFlag.Bool("up", false, "setup DB table")
		down := migrationFlag.Bool("down", false, "setup DB table")
		migrationFlag.Parse(os.Args[2:])
		return *up, *down
	}
	return false, false
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mahmad-labib/TodoInGo/db"
)

func main() {
	up, down := migrationFlag()
	conn, msg, err := db.InitDB()
	if err != nil {
		conn.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()
	if up {
		result, err := db.MigrateUp(conn)
		handleErr(err)
		msg = append(msg, result)
	}
	if down {
		result, err := db.MigrateDown(conn)
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
	migrationFlag := flag.NewFlagSet("migration", flag.ExitOnError)
	up := migrationFlag.Bool("up", false, "setup DB table")
	down := migrationFlag.Bool("down", false, "setup DB table")
	migrationFlag.Parse(os.Args[2:])
	return *up, *down
}

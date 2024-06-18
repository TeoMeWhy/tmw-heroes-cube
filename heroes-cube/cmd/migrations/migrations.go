package main

import (
	"heroes-cube/internals/db"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	con, err := db.Connect()
	if err != nil {
		log.Panicln(err)
	}

	db.MigrateClasses(con)
	db.MigrateRaces(con)

}

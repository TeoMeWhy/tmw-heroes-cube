package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "../data/database.db")
	return conn, err
}

func ImportQuery(path string) (*string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	query := string(b)

	return &query, nil
}

func ExecQueries(query string, con *sql.DB) error {

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	defer tx.Commit()

	queries := strings.Split(query, ";")
	for _, q := range queries {
		_, err := tx.Exec(q)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	}

	return nil

}

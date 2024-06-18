package db

import (
	"database/sql"
	"log"
)

type Class struct {
	Class           string
	PrimaryStatus   string
	SecondaryStatus string
	ThirdyStatus    string
}

func MigrateClasses(con *sql.DB) error {

	query, err := ImportQuery("internals/db/classes.sql")
	if err != nil {
		log.Println("Erro ao buscar query de construção de classes.", err)
		return err
	}
	return ExecQueries(*query, con)
}

func GetClass(class string, con *sql.DB) (*Class, error) {

	query := `
	SELECT
		Class,
		PrimaryStatus,
		SecondaryStatus,
		ThirdyStatus
	FROM classes
	WHERE Class = ?	
	`

	state, err := con.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := state.Query(class)
	if err != nil {
		return nil, err
	}

	c := &Class{}
	for rows.Next() {
		rows.Scan(&c.Class, &c.PrimaryStatus, &c.SecondaryStatus, &c.ThirdyStatus)
	}

	return c, nil

}

func GetClassList(con *sql.DB) (map[string]Class, error) {
	query := `
	SELECT
		Class,
		PrimaryStatus,
		SecondaryStatus,
		ThirdyStatus
	FROM classes
	`
	rows, err := con.Query(query)
	if err != nil {
		return nil, err
	}

	values := map[string]Class{}
	var c, p, s, t string
	for rows.Next() {

		rows.Scan(&c, &p, &s, &t)
		values[c] = Class{Class: c, PrimaryStatus: p, SecondaryStatus: s, ThirdyStatus: t}
	}

	return values, nil
}

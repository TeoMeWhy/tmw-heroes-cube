package db

import (
	"database/sql"
	"log"
)

type Race struct {
	Race        string
	Strength    int
	Agility     int
	Inteligence int
}

func MigrateRaces(con *sql.DB) error {

	query, err := ImportQuery("internals/db/races.sql")
	if err != nil {
		log.Println("Erro ao buscar query de construção de raças.", err)
		return err
	}

	return ExecQueries(*query, con)

}

func GetRace(race string, con *sql.DB) (*Race, error) {

	query := `
	SELECT 
		Race,
		Strength,
		Agility,
		Inteligence

	FROM races
	WHERE Race = ?`

	state, err := con.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := state.Query(race)
	if err != nil {
		return nil, err
	}

	r := &Race{}
	for rows.Next() {
		rows.Scan(
			&r.Race,
			&r.Strength,
			&r.Agility,
			&r.Inteligence,
		)
	}

	return r, nil

}

func GetRaceList(con *sql.DB) (map[string]Race, error) {

	query := `
	SELECT 
		Race,
		Strength,
		Agility,
		Inteligence

	FROM races`

	rows, err := con.Query(query)
	if err != nil {
		return nil, err
	}

	var r string
	var s, a, i int
	values := map[string]Race{}

	for rows.Next() {
		rows.Scan(&r, &s, &a, &i)
		values[r] = Race{Race: r, Strength: s, Agility: a, Inteligence: i}
	}

	return values, nil
}

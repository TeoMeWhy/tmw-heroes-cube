package db

import "database/sql"

type Item struct {
	Id          string
	Name        string
	Weight      int
	Strength    int
	Agility     int
	Inteligence int
	Damage      int
	HitPoints   int
	Defense     int
	Type        string
}

func GetItem(idItem string, con *sql.DB) (*Item, error) {

	query := `
	SELECT
		Id,
		Name,
		Weight,
		Strength,
		Agility,
		Inteligence,
		Damage,
		HitPoints,
		Defense,
		Type
	FROM items
	WHERE Id = ?		
	`

	state, err := con.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := state.Query(idItem)
	if err != nil {
		return nil, err
	}

	item := &Item{}
	for rows.Next() {
		rows.Scan(
			&item.Id,
			&item.Name,
			&item.Weight,
			&item.Strength,
			&item.Agility,
			&item.Inteligence,
			&item.Damage,
			&item.HitPoints,
			&item.Defense,
			&item.Type,
		)
	}

	return item, nil
}

func GetItemList(con *sql.DB) ([]Item, error) {

	query := `
	SELECT 
		Id,
		Name,
		Weight,
		Strength,
		Agility,
		Inteligence,
		Damage,
		HitPoints,
		Defense,
		Type
	
	FROM items`

	rows, err := con.Query(query)
	if err != nil {
		return nil, err
	}

	items := []Item{}
	for rows.Next() {
		i := Item{}
		rows.Scan(
			&i.Id,
			&i.Name,
			&i.Weight,
			&i.Strength,
			&i.Agility,
			&i.Inteligence,
			&i.Damage,
			&i.HitPoints,
			&i.Defense,
			&i.Type,
		)
		items = append(items, i)
	}

	return items, nil

}

package db

import (
	"database/sql"
)

type InventoryItem struct {
	IdPerson string
	IdItem   string
	Quantity int
}

type Inventory []InventoryItem

func CreateInventory(i Inventory, con *sql.DB) error {

	tx, err := con.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	for _, item := range i {
		if err := CreateInventoryItem(item, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

func CreateInventoryItem(item InventoryItem, tx *sql.Tx) error {

	query := `
	INSERT INTO inventories (IdPerson,IdItem, Quantity)
	VALUES (?,?,?)
	`

	state, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(item.IdPerson, item.IdItem, item.Quantity)
	if err != nil {
		return err
	}

	return err

}

func DeleteInventory(idPerson string, con *sql.DB) error {

	query := `DELETE FROM inventories
	WHERE IdPerson = ?`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(idPerson)
	return err
}

func DeleteInventoryItem(i InventoryItem, tx *sql.Tx) error {

	query := `DELETE FROM inventories
	WHERE IdPerson = ?
	AND IdItem = ?
	`

	state, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(i.IdPerson, i.IdItem)
	return err
}

func UpdateInventoryItem(i InventoryItem, con *sql.DB) error {

	query := `
	UPDATE inventories
	SET
		IdPerson=?,
		IdItem=?,
		Quantity=?

	WHERE IdPerson=?
	AND IdItem=?`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(i.IdPerson, i.IdItem, i.Quantity, i.IdPerson, i.IdItem)

	return err

}

func UpdateInventory(i Inventory, con *sql.DB) error {

	for _, item := range i {
		if err := UpdateInventoryItem(item, con); err != nil {
			return err
		}
	}

	return nil
}

func GetInventory(idPerson string, con *sql.DB) (Inventory, error) {

	query := `
	SELECT
    	IdPerson,
		IdItem,
		Quantity

	FROM inventories

	WHERE IdPerson = ?
	`

	state, err := con.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := state.Query(idPerson)
	if err != nil {
		return nil, err
	}

	inventory := []InventoryItem{}
	for rows.Next() {

		i := InventoryItem{}

		rows.Scan(
			&i.IdPerson,
			&i.IdItem,
			&i.Quantity,
		)

		inventory = append(inventory, i)
	}

	return inventory, nil
}

package db

import "database/sql"

type Slot struct {
	IdPerson string
	SlotPos  string
	IdItem   string
}

type Slots []Slot

func CreateSlot(s *Slot, con *sql.DB) error {

	query := `
	INSERT INTO slots
	(IdPerson,SlotPos,IdItem)
	VALUES (?,?,?)`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := state.Exec(
		s.IdPerson,
		s.SlotPos,
		s.IdItem,
	); err != nil {
		return err
	}

	return nil
}

func CreateSlots(slot *Slots, con *sql.DB) error {

	for _, s := range *slot {
		if err := CreateSlot(&s, con); err != nil {
			return err
		}
	}

	return nil
}

func GetSlots(idPerson string, con *sql.DB) (Slots, error) {
	query := `
	SELECT
		IdPerson,
		SlotPos,
		IdItem
	FROM slots
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

	slots := Slots{}
	for rows.Next() {
		s := Slot{}
		rows.Scan(
			&s.IdPerson,
			&s.SlotPos,
			&s.IdItem,
		)
		slots = append(slots, s)
	}

	return slots, nil
}

func DeleteSlots(idPerson string, con *sql.DB) error {

	query := `
	DELETE FROM slots
	WHERE IdPerson = ?
	`

	states, err := con.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := states.Exec(idPerson); err != nil {
		return err
	}

	return nil
}

func UpdateSlot(s *Slot, con *sql.DB) error {

	query := `
	UPDATE slots
	SET
		IdPerson = ?,
		SlotPos = ?,
		IdItem = ?
	
	WHERE IdPerson = ?
	AND SlotPos = ?
	`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := state.Exec(
		s.IdPerson,
		s.SlotPos,
		s.IdItem,
		s.IdPerson,
		s.SlotPos,
	); err != nil {
		return err
	}

	return nil
}

func UpdateSlots(slots *Slots, idPerson string, con *sql.DB) error {

	err := DeleteSlots(idPerson, con)
	if err != nil {
		return err
	}

	for _, s := range *slots {

		err := CreateSlot(&s, con)
		if err != nil {
			return err
		}
	}

	return nil
}

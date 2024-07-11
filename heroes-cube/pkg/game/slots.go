package game

import (
	"heroes-cube/internals/db"
	"log"
)

type Slots map[string]Item

func ImportSlots(idPerson string) (Slots, error) {

	slots := map[string]Item{}

	slotsDB, err := db.GetSlots(idPerson, con)
	if err != nil {
		return nil, err
	}

	for _, slotDB := range slotsDB {

		item, err := ImportItem(slotDB.IdItem)
		if err != nil {
			return nil, err
		}
		slots[slotDB.SlotPos] = *item
	}
	return slots, nil
}

func (s *Slots) ToSlotsDB(idPerson string) db.Slots {

	slotsDB := db.Slots{}

	for k, v := range *s {
		slot := db.Slot{
			IdPerson: idPerson,
			SlotPos:  k,
			IdItem:   v.Id,
		}
		slotsDB = append(slotsDB, slot)
	}
	return slotsDB
}

func (s *Slots) UpdateOrCreate(idPerson string) error {

	slotsDB := s.ToSlotsDB(idPerson)

	slotsDBActual, err := db.GetSlots(idPerson, con)
	if err != nil {
		return err
	}

	if len(slotsDBActual) == 0 {
		log.Println("Slots não existe, criando...")
		return db.CreateSlots(&slotsDB, con)
	}

	log.Println("Slots existe, atualizando...")
	return db.UpdateSlots(&slotsDB, idPerson, con)
}

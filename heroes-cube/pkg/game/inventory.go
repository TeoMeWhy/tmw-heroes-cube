package game

import "heroes-cube/internals/db"

type InventoryItem struct {
	Item
	Quantity int
}

type Inventory []InventoryItem

func ImportInventory(idPerson string) (Inventory, error) {

	inventory := Inventory{}

	inventoryDB, err := db.GetInventory(idPerson, con)
	if err != nil {
		return nil, err
	}

	for _, itemDB := range inventoryDB {

		idItem := itemDB.IdItem
		quantity := itemDB.Quantity

		item, err := ImportItem(idItem)
		if err != nil {
			return nil, err
		}

		inventory = append(inventory, InventoryItem{
			Item:     *item,
			Quantity: quantity,
		})

	}
	return inventory, nil
}

func (inventory Inventory) ToinventoryDB(idPerson string) db.Inventory {

	inventoryDB := db.Inventory{}

	for _, i := range inventory {
		inventoryItemDB := db.InventoryItem{
			IdPerson: idPerson,
			IdItem:   i.Id,
			Quantity: i.Quantity,
		}

		inventoryDB = append(inventoryDB, inventoryItemDB)
	}
	return inventoryDB
}

func (inventory Inventory) Save(idPerson string) error {

	inventoryDB := inventory.ToinventoryDB(idPerson)

	inventoryDBActual, err := db.GetInventory(idPerson, con)
	if err != nil {
		return err
	}

	if len(inventoryDBActual) == 0 {
		return db.CreateInventory(inventoryDB, con)
	}
	return db.UpdateInventory(inventoryDB, con)
}

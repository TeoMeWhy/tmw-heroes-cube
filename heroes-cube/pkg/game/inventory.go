package game

import (
	"heroes-cube/internals/db"
	"heroes-cube/internals/utils"
)

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

func (inventory Inventory) UpdateOrCreate(idPerson string) error {

	inventoryDB := inventory.ToinventoryDB(idPerson)
	if err := db.DeleteInventory(idPerson, con); err != nil {
		return err
	}

	if err := db.CreateInventory(inventoryDB, con); err != nil {
		return err
	}

	return nil
}

func (inventory Inventory) AddItem(idItem string) Inventory {

	item := Items[idItem]
	inventoryItem := InventoryItem{
		Item:     item,
		Quantity: 1,
	}

	inventory = append(inventory, inventoryItem)

	return inventory
}

func (inventory Inventory) RemoveItem(idItem string) (Inventory, error) {

	if check := inventory.HaveItem(idItem); !check {
		return inventory, utils.ItemNotFoundInInventory
	}

	newInventory := Inventory{}

	removed := false
	for _, i := range inventory {

		if i.Id == idItem && !removed {

			if i.Quantity > 1 {
				i.Quantity--
				newInventory = append(newInventory, i)
			}

			removed = true
			continue
		}
		newInventory = append(newInventory, i)
	}

	return newInventory, nil
}

func (inventory Inventory) HaveItem(idItem string) bool {
	for _, i := range inventory {
		if i.Id == idItem {
			return true
		}
	}
	return false
}

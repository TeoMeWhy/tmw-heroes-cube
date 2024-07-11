package handler

import (
	"heroes-cube/internals/utils"
	"heroes-cube/pkg/game"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETinventory(c *gin.Context) {

	param := c.Param("id")

	inventory, err := game.ImportInventory(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func PUTinventory(c *gin.Context) {

	type bodySchema struct {
		IdPlayer string `json:"idplayer" binding:"required"`
		IdItem   string `json:"iditem" binding:"required"`
		Op       string `json:"op" binding:"required"`
	}

	body := &bodySchema{}

	err := c.ShouldBind(body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if body.Op == "add" {
		err = InventoryAddItem(body.IdPlayer, body.IdItem)
	} else if body.Op == "remove" {
		err = InventoryRemoveItem(body.IdPlayer, body.IdItem)
	}

	if err == utils.PersonNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "inventário atualizado"})
}

func InventoryAddItem(idPlayer, idItem string) error {

	person, err := game.ImportPerson(idPlayer)
	if err != nil {
		return err
	}

	if person.Id != idPlayer {
		return utils.PersonNotFound
	}

	item := game.Items[idItem]
	inventoryItem := game.InventoryItem{
		Item:     item,
		Quantity: 1,
	}

	person.Inventory = append(person.Inventory, inventoryItem)

	if err := person.UpdateOrCreate(); err != nil {
		return err
	}

	return nil
}

func InventoryRemoveItem(idPlayer, idItem string) error {

	person, err := game.ImportPerson(idPlayer)
	if err != nil {
		return err
	}

	if person.Id != idPlayer {
		return utils.PersonNotFound
	}

	newInventory := game.Inventory{}

	removed := false
	for _, i := range person.Inventory {

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

	person.Inventory = newInventory

	if err := person.UpdateOrCreate(); err != nil {
		return err
	}

	return nil

}

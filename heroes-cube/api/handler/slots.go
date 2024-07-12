package handler

import (
	"heroes-cube/internals/utils"
	"heroes-cube/pkg/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETSlots(c *gin.Context) {

	id := c.Param("id")
	slots, err := game.ImportSlots(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if len(slots) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "inventário vazio ou não encontrado"})
		return
	}
	c.JSON(http.StatusOK, slots)
}

func PUTSlots(c *gin.Context) {

	type bodySchema struct {
		IdPlayer string `json:"idplayer" binding:"required"`
		IdItem   string `json:"iditem" binding:"required"`
		Op       string `json:"op" binding:"required"`
	}

	body := &bodySchema{}

	err := c.ShouldBind(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "utilize os campos necessários (idplayer, iditem, op)"})
		return
	}

	if body.Op == "add" {

		err := addSlots(body.IdPlayer, body.IdItem)
		if err == utils.ItemNotFoundInInventory {
			c.JSON(http.StatusBadRequest, gin.H{"error": "você não tem esse item no seu inventário"})
			return
		} else if err == utils.ItemNotCompatible {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sua classe não pode equipar esse item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})

	} else if body.Op == "remove" {
		return
	}

}

func addSlots(idPlayer, idItem string) error {

	newItem := game.Items[idItem]

	person, err := game.ImportPerson(idPlayer)
	if err != nil {
		return err
	}

	if check := person.Inventory.HaveItem(idItem); !check {
		return utils.ItemNotFoundInInventory
	}

	oldItem, ok := person.Slots[newItem.Type]
	if ok {
		person.Inventory = person.Inventory.AddItem(oldItem.Id)
	}

	err = person.EquipItem(idItem)
	if err != nil {
		return err
	}
	person.Inventory, err = person.Inventory.RemoveItem(idItem)
	if err != nil {
		return err
	}

	person.UpdateOrCreate()
	return nil

}

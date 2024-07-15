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
		return

	} else if body.Op == "remove" {

		err := removeSlots(body.IdPlayer, body.IdItem)
		if err == utils.ItemNotFoundInSlot {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return

		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sua classe não pode equipar esse item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Operação inválida, utilize 'add' ou 'remove'"})

}

func addSlots(idPlayer, idItem string) error {

	person, err := game.ImportPerson(idPlayer)
	if err != nil {
		return err
	}

	err = person.EquipItem(idItem)
	if err != nil {
		return err
	}

	person.UpdateOrCreate()
	return nil
}

func removeSlots(idPlayer, idItem string) error {

	item := game.Items[idItem]

	person, err := game.ImportPerson(idPlayer)
	if err != nil {
		return err
	}

	err = person.UnequipItem(item.Type)
	if err != nil {
		return err
	}

	person.UpdateOrCreate()
	return nil
}

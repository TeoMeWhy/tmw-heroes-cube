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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	inventory, err := game.ImportInventory(body.IdPlayer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if body.Op == "add" {

		inventory = inventory.AddItem(body.IdItem)

	} else if body.Op == "remove" {

		inventory, err = inventory.RemoveItem(body.IdItem)
		if err == utils.ItemNotFoundInInventory {
			c.JSON(http.StatusBadGateway, gin.H{"status": "item não encontrado no inventário"})
			return
		}
	}

	if err := inventory.UpdateOrCreate(body.IdPlayer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "inventário atualizado"})
}

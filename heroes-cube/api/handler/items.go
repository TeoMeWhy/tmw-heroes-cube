package handler

import (
	"heroes-cube/pkg/game"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETitems(c *gin.Context) {
	c.JSON(http.StatusOK, game.Items)
}

func GETrandomItem(c *gin.Context) {

	item, err := game.SelectRandomItem(game.Items)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, item)
}

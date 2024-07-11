package handler

import (
	"heroes-cube/pkg/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETraces(c *gin.Context) {
	c.JSON(http.StatusOK, game.Races)
}

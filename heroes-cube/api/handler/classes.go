package handler

import (
	"heroes-cube/pkg/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETclasses(c *gin.Context) {

	c.JSON(http.StatusOK, game.Classes)

}

package main

import (
	"heroes-cube/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(gin.Recovery())

	r.POST("/persons", handler.POSTpersons)
	r.GET("/persons/:id", handler.GETpersons)

	r.GET("/inventories/:id", handler.GETinventory)
	r.PUT("/inventories/", handler.PUTinventory)

	r.GET("/races", handler.GETraces)
	r.GET("/classes", handler.GETclasses)
	r.GET("/items", handler.GETitems)
	r.GET("/random_item", handler.GETrandomItem)

	r.Run("localhost:8085")

}

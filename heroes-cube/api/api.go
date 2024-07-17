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
	r.GET("/persons/", handler.GETpersonsByName)

	r.GET("/inventories/:id", handler.GETinventory)
	r.PUT("/inventories/", handler.PUTinventory)

	r.GET("/races", handler.GETraces)
	r.GET("/classes", handler.GETclasses)
	r.GET("/items", handler.GETitems)
	r.GET("/items/:id", handler.GETitem)
	r.GET("/random_item", handler.GETrandomItem)

	r.GET("/slots/:id", handler.GETSlots)
	r.PUT("/slots/", handler.PUTSlots)

	r.Run("0.0.0.0:8085")

}

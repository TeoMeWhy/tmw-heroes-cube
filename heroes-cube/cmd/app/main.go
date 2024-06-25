package main

import (
	"heroes-cube/pkg/game"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	teo, err := game.NewPerson("1", "Téo Calvo", "mage", "human")
	if err != nil {
		log.Panic(err)
	}

	teo.Slots["head"] = game.Items["3"]
	teo.Slots["hands"] = game.Items["1"]

	if err := teo.Save(); err != nil {
		log.Panic(err)
	}

	log.Println(teo)
}

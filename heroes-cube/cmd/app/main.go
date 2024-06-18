package main

import (
	"heroes-cube/internals/db"
	"heroes-cube/pkg/game"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// p, _ := game.NewPerson("1234", "Téo Calvo", "Mage", "Human")
	// log.Println(p)

	// err := p.Create()
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println("Usuário criado com sucesso!")

	con, _ := db.Connect()
	dbp, err := db.GetPerson("1234", con)
	if err != nil {
		log.Println(err)
	}
	p := game.DBPersonToPerson(dbp)

	p.Skills["Strength"] += 100

	if err := p.Save(); err != nil {
		log.Panicln("Erro ao salvar")
	}

	log.Println("Sucesso ao salvar")

}

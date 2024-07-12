package handler

import (
	"errors"
	"heroes-cube/internals/utils"
	"heroes-cube/pkg/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonBody struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Race        string `json:"race" binding:"required"`
	Class       string `json:"class" binding:"required"`
	Strength    int    `json:"strength"`
	Agility     int    `json:"agility"`
	Inteligence int    `json:"inteligence"`
	Damage      int    `json:"damage"`
	HitPoints   int    `json:"hitPoints"`
	Defense     int    `json:"defense"`
	Exp         int    `json:"exp"`
	Level       int    `json:"level"`
}

func POSTpersons(c *gin.Context) {

	newPerson := &PersonBody{}

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := game.NewPerson(newPerson.ID, newPerson.Name, newPerson.Class, newPerson.Race)

	if errors.Is(err, utils.ClassNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, utils.RaceNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := p.UpdateOrCreate(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})

}

func GETpersons(c *gin.Context) {

	idParam := c.Param("id")

	body, err := importPersonById(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, body)
}

func GETpersonsByName(c *gin.Context) {
	name := c.Query("name")

	body, err := importPersonByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, body)

}

func importPersonByName(name string) (*PersonBody, error) {

	person, err := game.ImportPersonbyName(name)
	if err != nil {
		return nil, err
	}

	if person.Id == "" {
		return nil, utils.PersonNotFound
	}

	pDB := person.ToPersonDB()

	bodyPerson := &PersonBody{
		ID:          pDB.Id,
		Name:        pDB.Name,
		Race:        pDB.Race,
		Class:       pDB.Class,
		Strength:    pDB.Strength,
		Agility:     pDB.Agility,
		Inteligence: pDB.Inteligence,
		Damage:      pDB.Damage,
		HitPoints:   pDB.HitPoints,
		Defense:     pDB.Defense,
		Exp:         pDB.Exp,
		Level:       pDB.Level,
	}

	return bodyPerson, nil
}

func importPersonById(id string) (*PersonBody, error) {

	person, err := game.ImportPerson(id)
	if err != nil {
		return nil, err
	}

	if person.Id == "" {
		return nil, utils.PersonNotFound
	}

	pDB := person.ToPersonDB()

	bodyPerson := &PersonBody{
		ID:          pDB.Id,
		Name:        pDB.Name,
		Race:        pDB.Race,
		Class:       pDB.Class,
		Strength:    pDB.Strength,
		Agility:     pDB.Agility,
		Inteligence: pDB.Inteligence,
		Damage:      pDB.Damage,
		HitPoints:   pDB.HitPoints,
		Defense:     pDB.Defense,
		Exp:         pDB.Exp,
		Level:       pDB.Level,
	}

	return bodyPerson, nil

}

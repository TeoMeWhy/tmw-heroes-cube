package game

import (
	"heroes-cube/internals/db"
	"heroes-cube/internals/utils"
	"log"
	"sort"
)

type Person struct {
	Id        string
	Name      string
	Skills    map[string]int
	Damage    int
	HitPoints int
	Defense   int
	Race
	Class
	Slots
	Inventory
	Exp   int
	Level int
}

func (p *Person) initSkills() error {

	numbers := utils.DrawDM(12, 5)
	sort.IntSlice.Sort(numbers)
	numbers = numbers[1 : len(numbers)-1]

	skills := map[string]int{
		p.PrimaryStatus:   numbers[2],
		p.SecondaryStatus: numbers[1],
		p.ThirdyStatus:    numbers[0],
	}

	p.Skills["Agility"] = skills["Agility"]
	p.Skills["Inteligence"] = skills["Inteligence"]
	p.Skills["Strength"] = skills["Strength"]

	return nil
}

func (p *Person) initHitPoints() {
	p.HitPoints = p.Skills["Strength"] + 10
}

func (p *Person) initDamage() {
	p.Damage = p.Skills[p.PrimaryStatus]
	if p.Slots.Arms.Type == "weapon" {
		p.Damage += p.Slots.Arms.Damage
	}
}

func (p *Person) initDefense() {
	p.Defense = p.Skills["Agility"]

	if p.Slots.Arms.Type == "armor" {
		p.Damage += p.Slots.Arms.Defense
	}

	if p.Slots.Head.Type == "armor" {
		p.Damage += p.Slots.Head.Defense
	}

	if p.Slots.Chest.Type == "armor" {
		p.Damage += p.Slots.Chest.Defense
	}

	if p.Slots.Legs.Type == "armor" {
		p.Damage += p.Slots.Legs.Defense
	}
}

func (p *Person) Create() error {
	dbPerson := PersonToDBPerson(p)
	return db.CreatePerson(dbPerson, con)
}

func (p *Person) Save() error {
	dbPerson := PersonToDBPerson(p)
	return db.UpdatePerson(dbPerson, con)
}

func (p *Person) Delete() error {
	dbPerson := PersonToDBPerson(p)
	return db.DeletePerson(dbPerson, con)
}

func NewPerson(id, name, class, race string) (*Person, error) {

	pRace := Race{
		Race: dbRaces[race].Race,
		Modifiers: map[string]int{
			"Strength":    dbRaces[race].Strength,
			"Inteligence": dbRaces[race].Inteligence,
			"Agility":     dbRaces[race].Agility},
	}

	pClass := Class{
		Class:           dbClasses[class].Class,
		PrimaryStatus:   dbClasses[class].PrimaryStatus,
		SecondaryStatus: dbClasses[class].SecondaryStatus,
		ThirdyStatus:    dbClasses[class].ThirdyStatus,
	}

	p := &Person{
		Id:        id,
		Name:      name,
		Skills:    map[string]int{},
		Damage:    0,
		HitPoints: 0,
		Defense:   0,
		Race:      pRace,
		Class:     pClass,
		Slots:     Slots{},
		Inventory: []Item{},
		Exp:       0,
		Level:     0,
	}

	err := p.initSkills()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	p.initHitPoints()
	p.initDamage()
	p.initDefense()

	return p, nil

}

func PersonToDBPerson(p *Person) *db.Person {
	return &db.Person{
		Id:          p.Id,
		Name:        p.Name,
		Class:       p.Class.Class,
		Strength:    p.Skills["Strength"],
		Agility:     p.Skills["Agility"],
		Inteligence: p.Skills["Inteligence"],
		Damage:      p.Damage,
		HitPoints:   p.HitPoints,
		Defense:     p.Defense,
		Race:        p.Race.Race,
		Exp:         p.Exp,
		Level:       p.Level,
	}
}

func DBPersonToPerson(dbp *db.Person) *Person {

	class := Class{
		Class:           dbp.Class,
		PrimaryStatus:   dbClasses[dbp.Class].PrimaryStatus,
		SecondaryStatus: dbClasses[dbp.Class].SecondaryStatus,
		ThirdyStatus:    dbClasses[dbp.Class].ThirdyStatus,
	}

	skills := map[string]int{
		"Strength":    dbp.Strength,
		"Agility":     dbp.Agility,
		"Inteligence": dbp.Inteligence,
	}

	race := Race{
		Race: dbp.Race,
		Modifiers: map[string]int{
			"Strength":    dbRaces[dbp.Race].Strength,
			"Agility":     dbRaces[dbp.Race].Agility,
			"Inteligence": dbRaces[dbp.Race].Inteligence,
		},
	}

	p := &Person{
		Id:        dbp.Id,
		Name:      dbp.Name,
		Class:     class,
		Skills:    skills,
		Damage:    dbp.Damage,
		HitPoints: dbp.HitPoints,
		Defense:   dbp.Defense,
		Race:      race,
		Exp:       dbp.Exp,
		Level:     dbp.Level,
	}

	return p
}

package game

import (
	"heroes-cube/internals/db"
	"heroes-cube/internals/utils"
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

func (p *Person) initSkills() {

	numbers := utils.DrawDM(12, 5)
	sort.IntSlice.Sort(numbers)
	numbers = numbers[1 : len(numbers)-1]

	skills := map[string]int{
		p.PrimaryStatus:   numbers[2],
		p.SecondaryStatus: numbers[1],
		p.ThirdyStatus:    numbers[0],
	}

	p.Skills["agility"] = skills["agility"]
	p.Skills["inteligence"] = skills["inteligence"]
	p.Skills["strength"] = skills["strength"]
}

func (p *Person) SetHitPoints() {
	p.HitPoints = p.Skills["strength"] + 10
}

func (p *Person) SetDamage() {
	p.Damage = p.Skills[p.PrimaryStatus] + p.Level
	for _, v := range p.Slots {
		p.Damage += v.Damage
	}
}

func (p *Person) SetDefense() {
	p.Defense = p.Skills["agility"]

	for _, v := range p.Slots {
		p.Defense += v.Defense
	}
}

func (p *Person) ToPersonDB() *db.Person {

	personDB := &db.Person{

		Id:          p.Id,
		Name:        p.Name,
		Strength:    p.Skills["strength"],
		Agility:     p.Skills["agility"],
		Inteligence: p.Skills["inteligence"],
		Damage:      p.Damage,
		HitPoints:   p.HitPoints,
		Defense:     p.Defense,
		Class:       p.Class.Class,
		Race:        p.Race.Race,
		Exp:         p.Exp,
		Level:       p.Level,
	}
	return personDB
}

func (p *Person) UpdateOrCreate() error {

	p.SetDamage()
	p.SetDefense()
	p.SetHitPoints()

	if err := p.Inventory.UpdateOrCreate(p.Id); err != nil {
		return err
	}

	if err := p.Slots.UpdateOrCreate(p.Id); err != nil {
		return err
	}

	personDB := p.ToPersonDB()
	personDBActual, err := db.GetPerson(p.Id, con)
	if err != nil {
		return err
	}

	if personDBActual.Id != p.Id {
		return db.CreatePerson(personDB, con)
	}
	return db.UpdatePerson(personDB, con)
}

func NewPerson(id, name, class, race string) (*Person, error) {

	checkClass := false
	for k := range Classes {
		if class == k {
			checkClass = true
			break
		}
	}

	if !checkClass {
		return nil, utils.ClassNotFound
	}

	checkRace := false
	for k := range Races {
		if race == k {
			checkRace = true
			break
		}
	}

	if !checkRace {
		return nil, utils.RaceNotFound
	}

	p := &Person{
		Id:        id,
		Name:      name,
		Skills:    map[string]int{},
		Damage:    0,
		HitPoints: 0,
		Defense:   0,
		Race:      Races[race],
		Class:     Classes[class],
		Slots:     Slots{},
		Inventory: Inventory{},
		Exp:       1,
		Level:     1,
	}

	p.initSkills()
	p.SetHitPoints()
	p.SetDamage()
	p.SetDefense()

	return p, nil
}

func ImportPerson(idPerson string) (*Person, error) {

	personDB, err := db.GetPerson(idPerson, con)
	if err != nil {
		return nil, err
	}

	p := &Person{
		Id:        personDB.Id,
		Name:      personDB.Name,
		Damage:    personDB.Damage,
		HitPoints: personDB.HitPoints,
		Defense:   personDB.Defense,
		Exp:       personDB.Exp,
		Level:     personDB.Level,
	}

	p.Skills = map[string]int{
		"strength":    personDB.Strength,
		"agility":     personDB.Agility,
		"inteligence": personDB.Inteligence,
	}

	p.Race = Races[personDB.Race]
	p.Class = Classes[personDB.Class]

	slots, err := ImportSlots(idPerson)
	if err != nil {
		return nil, err
	}
	p.Slots = slots

	inventory, err := ImportInventory(idPerson)
	if err != nil {
		return nil, err
	}
	p.Inventory = inventory

	return p, nil

}

func ImportPersonbyName(name string) (*Person, error) {

	id, err := db.GetPersonIDbyName(name, con)
	if err != nil {
		return nil, err
	}

	return ImportPerson(id)
}

func (p *Person) EquipItem(idItem string) error {

	newItem := Items[idItem]

	if check := p.Inventory.HaveItem(newItem); !check {
		return utils.ItemNotFoundInInventory
	}

	if newItem.Class != "all" && newItem.Class != p.Class.Class {
		return utils.ItemNotCompatible
	}

	oldItem, ok := p.Slots[newItem.Type]
	if ok {
		p.Inventory = p.Inventory.AddItem(oldItem)
	}

	p.Slots = p.Slots.AddItem(newItem)

	var err error
	p.Inventory, err = p.Inventory.RemoveItem(newItem)
	if err != nil {
		return err
	}

	return nil

}

func (p *Person) UnequipItem(pos string) error {

	item, ok := p.Slots[pos]
	if !ok {
		return utils.ItemNotFoundInSlot
	}

	p.Inventory = p.Inventory.AddItem(item)
	p.Slots = p.Slots.RemoveItem(item)
	return nil

}

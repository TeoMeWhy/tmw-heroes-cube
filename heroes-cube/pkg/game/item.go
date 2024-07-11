package game

import (
	"heroes-cube/internals/db"
	"strconv"

	"math/rand"
)

type Item struct {
	Id          string
	Name        string
	Weight      int
	Strength    int
	Agility     int
	Inteligence int
	Damage      int
	HitPoints   int
	Defense     int
	Type        string
	Price       int
	Class       string
}

func ItemDBToItem(i db.Item) *Item {
	item := &Item{
		Id:          i.Id,
		Name:        i.Name,
		Weight:      i.Weight,
		Strength:    i.Strength,
		Agility:     i.Agility,
		Inteligence: i.Inteligence,
		Damage:      i.Damage,
		HitPoints:   i.HitPoints,
		Defense:     i.Defense,
		Type:        i.Type,
		Price:       i.Price,
		Class:       i.Class,
	}

	return item
}

func ImportItem(itemID string) (*Item, error) {

	itemDB, err := db.GetItem(itemID, con)
	if err != nil {
		return nil, err
	}

	i := ItemDBToItem(*itemDB)
	return i, nil

}

func ImportItems() (map[string]Item, error) {

	itemsDB, err := db.GetItemList(con)
	if err != nil {
		return nil, err
	}

	items := map[string]Item{}
	for _, i := range itemsDB {

		items[i.Id] = *ItemDBToItem(i)

	}

	return items, nil
}

func SelectRandomItem(items map[string]Item) (*Item, error) {

	var maxId int64

	for i := range items {

		iInt, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}

		if maxId < int64(iInt) {
			maxId = int64(iInt)
		}
	}

	idSelected := rand.Int63n(maxId) + 1

	item := items[strconv.Itoa(int(idSelected))]
	return &item, nil
}

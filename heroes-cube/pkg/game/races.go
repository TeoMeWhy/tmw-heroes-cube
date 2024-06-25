package game

import "heroes-cube/internals/db"

type Race struct {
	Race      string
	Modifiers map[string]int
}

func raceDBtoRace(raceDB db.Race) *Race {

	race := &Race{
		Race: raceDB.Race,
		Modifiers: map[string]int{
			"strength":    raceDB.Strength,
			"agility":     raceDB.Agility,
			"inteligence": raceDB.Inteligence,
		},
	}
	return race
}

func ImportRace(raceName string) (*Race, error) {

	raceDB, err := db.GetRace(raceName, con)
	if err != nil {
		return nil, err
	}

	race := raceDBtoRace(*raceDB)
	return race, nil

}

func ImportRaces() (map[string]Race, error) {

	racesDB, err := db.GetRaceList(con)
	if err != nil {
		return nil, err
	}

	races := map[string]Race{}
	for k, v := range racesDB {
		races[k] = *raceDBtoRace(v)
	}

	return races, nil

}

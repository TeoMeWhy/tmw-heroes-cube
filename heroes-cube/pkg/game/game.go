package game

import (
	"heroes-cube/internals/db"
)

var con, _ = db.Connect()
var dbClasses, _ = db.GetClassList(con)
var dbRaces, _ = db.GetRaceList(con)

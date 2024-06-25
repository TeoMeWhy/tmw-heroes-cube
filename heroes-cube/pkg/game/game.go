package game

import (
	"heroes-cube/internals/db"
)

var con, _ = db.Connect()
var Classes, _ = ImportClasses()
var Races, _ = ImportRaces()
var Items, _ = ImportItems()

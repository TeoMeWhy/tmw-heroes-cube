DROP TABLE IF EXISTS persons;
CREATE TABLE IF NOT EXISTS persons (
	Id          VARCHAR(150),
	Name        VARCHAR(150),
	Strength    INTEGER,
	Agility     INTEGER,
	Inteligence INTEGER,
	Damage      INTEGER,
	HitPoints   INTEGER,
	Defense     INTEGER,
	Class       VARCHAR(50),
	Race        VARCHAR(50),
	Exp         INTEGER,
	Level       INTEGER
);
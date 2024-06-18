DROP TABLE IF EXISTS races;

CREATE TABLE IF NOT EXISTS races(
    Race VARCHAR(50),
    Strength INTEGER,
    Agility INTEGER,
    Inteligence INTEGER
);

INSERT INTO races VALUES
    ('Human', 1,1, 1),
    ('Elf', 0,2, 1),
    ('Dwarf', 3,0, 0),
    ('Hobbit', 0,3, 0),
    ('Poney', 0,0, 0);
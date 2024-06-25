DROP TABLE IF EXISTS races;

CREATE TABLE IF NOT EXISTS races(
    Race VARCHAR(50),
    Strength INTEGER,
    Agility INTEGER,
    Inteligence INTEGER
);

INSERT INTO races VALUES
    ('human', 1,1, 1),
    ('elf', 0,2, 1),
    ('dwarf', 3,0, 0),
    ('hobbit', 0,3, 0),
    ('poney', 0,0, 0);
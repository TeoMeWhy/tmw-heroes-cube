DROP TABLE IF EXISTS items;
CREATE TABLE IF NOT EXISTS items (
    Id          VARCHAR(50),
    Name        VARCHAR(50),
    Weight      INTEGER,
    Strength    INTEGER,
    Agility     INTEGER,
    Inteligence INTEGER,
    Damage      INTEGER,
    HitPoints   INTEGER,
    Defense     INTEGER,
    Type        VARCHAR(50)
);

INSERT INTO  items VALUES 
    (1,'Espada do Dragão',5000, 10, 2,0,15,0,0,'weapon'),
    (2,'Escudo do Guardião',8000, 5, -2,0,0,0,20,'armor'),
    (3,'Elmo da Sabedoria',3000,0,0, 8,0,5,5,'armor'),
    (4,'Botas da Velocidade',1000,0, 10,0,0,0,2,'armor'),
    (5,'Cajado do Arcanjo',4000,0,0, 15,10,0,0,'weapon'),
    (6,'Armadura do Berserker',12000, 8, -4,0,0,20,15,'armor'),
    (7,'Anel da Vida',100,0,0, 5,0,10,0,'armor'),
    (8,'Adaga das Sombras',2000, 5, 8,0,12,0,0,'weapon'),
    (9,'Capa da Invisibilidade',500,0,5,2,0,0,10,'armor'),
    (10,'Poção de Cura',200,0,0,0,0,20,0,'utility')
;
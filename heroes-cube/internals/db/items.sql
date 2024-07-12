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
    Type        VARCHAR(50),
    Price       INTEGER,
    Class       VARCHAR(50)
);

INSERT INTO  items VALUES 
    (1,'Bota das Sombras',5,2,8,3,2,0,5,'foot',200, 'all'),
    (2,'Botas de Couraça',15,5,-2,0,3,10,15,'foot',500, 'all'),
    (3,'Chinélos Místicos',3,0,4,7,1,5,3,'foot',350, 'mage'),
    (4,'Botas Berserker',10,8,4,0,5,0,8,'foot',450, 'warrior'),
    (5,'Botas Diversão',10,8,4,0,10,0,8,'foot',450, 'bard'),
    (6,'Chapéu da Sombra',2,1,6,4,1,0,3,'head',150, 'all'),
    (7,'Capacete de Ferro',8,4,-1,0,3,8,12,'head',400, 'all'),
    (8,'Chapéu Arcano',1,0,3,9,1,4,2,'head',300, 'mage'),
    (9,'Elmo do Berserker',6,7,2,0,4,0,6,'head',350, 'warrior'),
    (10,'Couraça das Sombras',12,3,6,2,3,0,10,'chest',300, 'all'),
    (11,'Armadura de Placas',30,7,-3,0,4,15,25,'chest',700, 'all'),
    (12,'Túnica Arcana',5,0,4,10,1,7,5,'chest',500, 'mage'),
    (13,'Couraça do Berserker',20,10,3,0,5,0,12,'chest',600, 'warrior'),
    (14,'Espada Curta',5,3,2,0,5,0,0,'weapon',150, 'all'),
    (15,'Machado de Batalha',12,6,-1,0,7,0,2,'weapon',300, 'all'),
    (16,'Adaga',2,1,8,0,3,0,0,'weapon',100, 'all'),
    (17,'Martelo de Guerra',15,8,-2,0,7,0,5,'weapon',400, 'all'),
    (18,'Cajado Arcano',4,0,2,10,10,0,1,'weapon',250, 'mage'),
    (19,'Arco Longo',6,2,7,0,11,0,0,'weapon',200, 'all'),
    (20,'Lança',8,5,3,0,12,0,2,'weapon',220, 'all'),
    (21,'Maça',10,6,1,0,9,0,3,'weapon',280, 'all'),
    (22,'Espada Longa',7,5,2,0,12,0,1,'weapon',250, 'warrior'),
    (23,'Adaga Envenenada',3,1,7,0,6,0,0,'weapon',150, 'thief'),
    (24,'Cajado Sagrado',4,0,2,10,10,0,1,'weapon',250, 'cleric'),
    (25,'Ukulele das Trevas',4,0,2,10,7,0,1,'weapon',250, 'bard')
;

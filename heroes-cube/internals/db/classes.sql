DROP TABLE IF EXISTS classes;

CREATE TABLE IF NOT EXISTS classes (
    Class VARCHAR(50),
    PrimaryStatus VARCHAR(50),
    SecondaryStatus VARCHAR(50),
    ThirdyStatus VARCHAR(50)
);

INSERT INTO classes VALUES
    ('mage', 'inteligence', 'agility', 'strength'),
    ('cleric', 'inteligence', 'strength', 'agility'),
    ('thief', 'agility', 'inteligence', 'strength'),
    ('warrior', 'strength', 'agility', 'inteligence'),
    ('bard', 'agility', 'strength', 'inteligence');

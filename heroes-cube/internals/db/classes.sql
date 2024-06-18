DROP TABLE IF EXISTS classes;

CREATE TABLE IF NOT EXISTS classes (
    Class VARCHAR(50),
    PrimaryStatus VARCHAR(50),
    SecondaryStatus VARCHAR(50),
    ThirdyStatus VARCHAR(50)
);

INSERT INTO classes VALUES
    ('Mage', 'Inteligence', 'Agility', 'Strength'),
    ('Cleric', 'Inteligence', 'Strength', 'Agility'),
    ('Thief', 'Agility', 'Inteligence', 'Strength'),
    ('Warrior', 'Strength', 'Agility', 'Inteligence'),
    ('Bard', 'Agility', 'Strength', 'Inteligence');

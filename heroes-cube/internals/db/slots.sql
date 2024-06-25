DROP TABLE IF EXISTS slots;
CREATE TABLE IF NOT EXISTS slots(
	IdPerson VARCHAR(50),
	SlotPos  VARCHAR(50), -- head, chest, arms, hands, legs, feet, back
	IdItem   VARCHAR(50)
);

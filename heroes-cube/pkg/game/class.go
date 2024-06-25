package game

import "heroes-cube/internals/db"

type Class struct {
	Class           string
	PrimaryStatus   string
	SecondaryStatus string
	ThirdyStatus    string
}

func ClassDBtoClass(classDB db.Class) *Class {

	c := &Class{
		Class:           classDB.Class,
		PrimaryStatus:   classDB.PrimaryStatus,
		SecondaryStatus: classDB.SecondaryStatus,
		ThirdyStatus:    classDB.ThirdyStatus,
	}

	return c

}

func ImportClass(className string) (*Class, error) {
	classDB, err := db.GetClass(className, con)
	if err != nil {
		return nil, err
	}

	c := ClassDBtoClass(*classDB)
	return c, nil

}

func ImportClasses() (map[string]Class, error) {

	classesDB, err := db.GetClassList(con)
	if err != nil {
		return nil, err
	}

	classes := map[string]Class{}
	for k, v := range classesDB {
		classes[k] = *ClassDBtoClass(v)
	}

	return classes, nil

}

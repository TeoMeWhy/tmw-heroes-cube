package utils

func CheckClass(class string, allClasses []string) bool {

	for _, v := range allClasses {
		if class == v {
			return true
		}
	}
	return false

}

func CheckRace(race string, allraces []string) bool {

	for _, v := range allraces {
		if race == v {
			return true
		}
	}
	return false

}

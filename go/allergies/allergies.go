package allergies

import "strings"

const testVersion = 1

var allergenScoreMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) (result []string) {
	for k, v := range allergenScoreMap {
		if score&k != 0 {
			result = append(result, v)
		}
	}
	return
}

func AllergicTo(score uint, allergen string) bool {
	allergens := Allergies(score)
	allergensString := strings.Join(allergens, " ")
	return strings.Contains(allergensString, allergen)
}

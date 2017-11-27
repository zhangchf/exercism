package raindrops

import (
	"strconv"
)

const testVersion = 3

func Convert(v int) string {
	var result string
	if v%3 == 0 {
		result += "Pling"
	}
	if v%5 == 0 {
		result += "Plang"
	}
	if v%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(v)
	}
	return result
}

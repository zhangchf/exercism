package etl

import (
	"strings"
)

const testVersion = 1

func Transform(old map[int][]string) (new map[string]int) {
	new = make(map[string]int)
	for key, value := range old {
		for _, c := range value {
			new[strings.ToLower(c)] = key
		}
	}
	return
}

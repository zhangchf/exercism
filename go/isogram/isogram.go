package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, r := range word {
		if unicode.IsLetter(r) && strings.ContainsRune(word[i+1:], r) {
			return false
		}
	}
	return true
}

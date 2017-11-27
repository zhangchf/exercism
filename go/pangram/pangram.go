package pangram

import (
	"strings"
	"unicode"
)

const testVersion = 2

func IsPangram(input string) bool {
	var chars = make(map[rune]bool)
	for _, c := range strings.ToLower(input) {
		if unicode.IsLetter(c) {
			chars[c] = true
		}
	}
	return len(chars) == 26
}

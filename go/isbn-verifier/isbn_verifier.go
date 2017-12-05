package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(input string) bool {
	input = strings.Map(func(r rune) rune {
		if r == '-' {
			return -1
		}
		return r
	}, input)

	if len(input) != 10 {
		return false
	}
	sum := 0
	for i, c := range input {
		if i < 9 && !unicode.IsDigit(c) {
			return false
		}
		num := int(c - '0')
		if i == 9 {
			if !(unicode.IsDigit(c) || c == 'X') {
				return false
			}
			if c == 'X' {
				num = 10
			}
		}
		sum += num * (10 - i)
	}
	if sum%11 == 0 {
		return true
	}
	return false
}

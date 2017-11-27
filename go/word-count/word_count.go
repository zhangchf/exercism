package wordcount

import (
	"fmt"
	"strings"
	"unicode"
)

const testVersion = 3

type Frequency map[string]int

// WordCount returns occurence of each word in the given phrase.
func WordCount(phrase string) Frequency {
	fields := strings.FieldsFunc(phrase, func(r rune) bool {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'') {
			return true
		}
		return false
	})

	fmt.Println(phrase, fields)

	result := make(Frequency)
	for _, field := range fields {
		field = strings.Trim(field, "'")
		result[strings.ToLower(field)] += 1
	}
	return result
}

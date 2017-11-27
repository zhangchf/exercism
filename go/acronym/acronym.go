package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(input string) string {
	if i := strings.IndexRune(input, ':'); i != -1 {
		return input[0:i]
	}

	fields := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == ',' || r == '-'
	})

	result := ""
	for _, field := range fields {
		result += field[0:1]
	}
	return strings.ToUpper(result)
}

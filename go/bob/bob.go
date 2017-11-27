package bob

import (
	"strings"
	"regexp"
)

const testVersion = 3

func Hey(input string) string {
	input = strings.TrimSpace(input)
	switch {
	case isShouting(input):
		return "Whoa, chill out!"
	case isQuestion(input):
		return "Sure."
	case isNothing(input):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}

func isQuestion(input string) bool {
	return strings.HasSuffix(input, "?")
}

func isShouting(input string) bool {
	return hasLetter(input) && strings.Compare(strings.ToUpper(input), input) == 0
}

func isNothing(input string) bool {
	return len(strings.Trim(input, " \n\r\t")) == 0
}

func hasLetter(input string) bool {
	matched, _ := regexp.MatchString(`.*[a-zA-Z].*`, input)
	return matched
}

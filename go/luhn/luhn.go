package luhn

import (
	"strings"
	"strconv"
)

const testVersion = 2

func Valid(input string) (valid bool) {
	input = strings.TrimSpace(input)

	if len(input) <= 1 {
		return false
	}

	sum := 0
	pos := 0
	for i := len(input)-1; i >= 0; i-- {
		if input[i] == ' ' {
			continue
		}
		num, err := strconv.Atoi(input[i:i+1])
		if err != nil {
			return false
		}
		pos++
		if pos%2 == 0 {
			num = num*2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	if sum % 10 == 0 {
		return true
	}
	return false
}

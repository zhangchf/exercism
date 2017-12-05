package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	output := ""
	previous := '1'
	count := 0
	for _, c := range input {
		if c != previous {
			if count != 0 {
				if count != 1 {
					output += strconv.Itoa(count)
				}
				output += string(previous)
			}
			previous = c
			count = 1
		} else {
			count++
		}
	}
	if count != 0 {
		if count != 1 {
			output += strconv.Itoa(count)
		}
		output += string(previous)
	}
	return output
}

func RunLengthDecode(encoded string) string {
	decoded := []byte{}
	numStr := ""
	for _, c := range encoded {
		if unicode.IsDigit(c) {
			numStr += string(c)
		} else {
			if numStr == "" {
				decoded = append(decoded, byte(c))
			} else if num, err := strconv.Atoi(numStr); err == nil {
				for i := 0; i < num; i++ {
					decoded = append(decoded, byte(c))
				}
			}
			numStr = ""
		}
	}
	return string(decoded)
}

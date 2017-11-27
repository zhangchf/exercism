package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

func Gen(char byte) (string, error) {
	if !(char >= 'A' && char <= 'Z') {
		return "", errors.New("Invalid input")
	}

	sideLen := char - 'A'
	rows := 2 * sideLen + 1

	result := make([]string, rows)

	row1 := make([]byte, rows)
	row2 := make([]byte, rows)
	for i := byte(0); i <= sideLen; i = i+1 {
		for m := byte(0); m <= sideLen; m = m+1 {
			c := byte(' ')
			if m == sideLen - i {
				c = 'A' + i
			}
			row1[m] = c
			row1[rows-1-m] = c
			result[i] = string(row1)
			if i != sideLen {
				row2[m] = c
				row2[rows-1-m] = c
				result[rows-1-i] = string(row2)
			}
		}
	}

	return strings.Join(result, "\n"), nil
}

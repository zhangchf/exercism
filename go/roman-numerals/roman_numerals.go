package romannumerals

import (
	"errors"
	"math"
	"strconv"
)

const testVersion = 4

var romans = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(arabic int) (romans string, err error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("Invalid input")
	}
	dstr := strconv.Itoa(arabic)

	dlen := len(dstr)
	for i := range dstr {
		d, err := strconv.Atoi(dstr[i : i+1])
		if err != nil {
			return "", errors.New("Invalid input")
		}
		romans += getRomanNumeral(d, int(math.Pow10(dlen-i-1)))
	}
	return
}

func getRomanNumeral(d, base int) (result string) {
	var r1, r5, r10 string
	switch base {
	case 1:
		r1, r5, r10 = romans[1], romans[5], romans[10]
	case 10:
		r1, r5, r10 = romans[10], romans[50], romans[100]
	case 100:
		r1, r5, r10 = romans[100], romans[500], romans[1000]
	case 1000:
		r1 = romans[1000]
	}

	switch d {
	case 1, 2, 3:
		for i := 0; i < d; i++ {
			result += r1
		}
	case 4:
		result = r1 + r5
	case 5:
		result = r5
	case 6, 7, 8:
		result = r5
		for i := 5; i < d; i++ {
			result += r1
		}
	case 9:
		result = r1 + r10
	}
	return
}

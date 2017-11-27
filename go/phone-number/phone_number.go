package phonenumber

import (
	"strings"
	"unicode"
	"errors"
)

var ErrInvalidInput = errors.New("Invalid Input")

func Number(in string) (string, error) {
	t := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		} else {
			return -1
		}
	}, in)

	tlen := len(t)
	if tlen == 10 {
		if !hasValidN(t) {
			return "", ErrInvalidInput
		}
		return t, nil
	} else if tlen == 11 {
		if !isValidCountryCode(t[0]) || !hasValidN(t[1:]) {
			return "", ErrInvalidInput
		}
		return t[1:], nil
	} else {
		return "", ErrInvalidInput
	}

}

func isValidN(r uint8) bool {
	if r >= '2' && r <= '9' {
		return true
	}
	return false
}

func isValidCountryCode(r uint8) bool {
	return r == '1'
}

func hasValidN(s string) bool {
	return isValidN(s[0]) && isValidN(s[3])
}

func AreaCode(in string) (string, error) {
	if num, err := Number(in); err != nil {
		return num, err
	} else {
		return num[:3], nil
	}
}

func Format(in string) (string, error) {
	if num, err := Number(in); err != nil {
		return num, err
	} else {
		return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
	}
}

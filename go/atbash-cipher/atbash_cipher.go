package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

func Atbash(in string) (out string) {
	bbuf := bytes.Buffer{}
	j := 0
	for _, v := range in {
		switch {
		case unicode.IsLetter(v):
			t := unicode.ToLower(v)
			bbuf.WriteRune('z' + 'a' - t)
		case unicode.IsDigit(v):
			bbuf.WriteRune(v)
		default:
			continue
		}
		j++
		if j > 0 && j%5 == 0 {
			bbuf.WriteRune(' ')
		}
	}
	temp := bbuf.String()
	return strings.TrimSpace(temp)
}

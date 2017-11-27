package igpay

import (
	"strings"
)

var vowels = "aeiou"
var consonant2 = map[string]bool{"ch":true, "th":true, "qu":true}
var consonant3 = map[string]bool{"squ":true, "thr":true, "sch":true}

func PigLatin(in string) (string) {
	parts := strings.Split(in, " ")
	out := make([]string, len(parts))
	for i := 0; i < len(parts); i++ {
		out[i] = plword(parts[i])
	}
	return strings.Join(out, " ")
}

func plword(in string) string {
	startPos := 0
	if strings.Contains(vowels, in[:1]) {
		startPos = 0
	} else if len(in) > 3 && consonant3[in[:3]] {
		startPos = 3
	} else if len(in) > 2 && consonant2[in[:2]] {
		startPos = 2
	} else {
		startPos = 1
	}
	return in[startPos:] + in[:startPos] + "ay"
}

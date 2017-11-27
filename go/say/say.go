package say

import "strings"

const testVersion = 1

const (
	HUNDRED     uint64 = 100
	THOUSAND           = 1000
	MILLION            = THOUSAND * THOUSAND
	BILLION            = MILLION * THOUSAND
	TRILLION           = BILLION * THOUSAND
	QUADRILLION        = TRILLION * THOUSAND
	QUINTILLION        = QUADRILLION * THOUSAND
)

var num2str = map[uint64]string{
	0:           "zero",
	1:           "one",
	2:           "two",
	3:           "three",
	4:           "four",
	5:           "five",
	6:           "six",
	7:           "seven",
	8:           "eight",
	9:           "nine",
	10:          "ten",
	11:          "eleven",
	12:          "twelve",
	13:          "thirteen",
	14:          "fourteen",
	15:          "fifteen",
	16:          "sixteen",
	17:          "seventeen",
	18:          "eighteen",
	19:          "nineteen",
	20:          "twenty",
	30:          "thirty",
	40:          "forty",
	50:          "fifty",
	60:          "sixty",
	70:          "seventy",
	80:          "eighty",
	90:          "ninety",
	HUNDRED:     "hundred",
	THOUSAND:    "thousand",
	MILLION:     "million",
	BILLION:     "billion",
	TRILLION:    "trillion",
	QUADRILLION: "quadrillion",
	QUINTILLION: "quintillion",
}

// Say transforms a number into a string. \n
// A more concise implementation: http://exercism.io/submissions/25aa351d60dd4136a40bc2b058b13a3a
func Say(num uint64) string {
	if num == 0 {
		return num2str[num]
	}

	quintillions, remainder := num/QUINTILLION, num%QUINTILLION
	quadrillions, remainder := remainder/QUADRILLION, remainder%QUADRILLION
	trillions, remainder := remainder/TRILLION, remainder%TRILLION
	billions, remainder := remainder/BILLION, remainder%BILLION
	millions, remainder := remainder/MILLION, remainder%MILLION
	thousands, remainder := remainder/THOUSAND, remainder%THOUSAND

	parts := make([]string, 0)
	if quintillions > 0 {
		parts = append(parts, say1to999(quintillions))
		parts = append(parts, num2str[QUINTILLION])
	}
	if quadrillions > 0 {
		parts = append(parts, say1to999(quadrillions))
		parts = append(parts, num2str[QUADRILLION])
	}
	if trillions > 0 {
		parts = append(parts, say1to999(trillions))
		parts = append(parts, num2str[TRILLION])
	}
	if billions > 0 {
		parts = append(parts, say1to999(billions))
		parts = append(parts, num2str[BILLION])
	}
	if millions > 0 {
		parts = append(parts, say1to999(millions))
		parts = append(parts, num2str[MILLION])
	}
	if thousands > 0 {
		parts = append(parts, say1to999(thousands))
		parts = append(parts, num2str[THOUSAND])
	}
	if remainder > 0 {
		parts = append(parts, say1to999(remainder))
	}

	return strings.Join(parts, " ")

}

func say0to99(num uint64) string {
	if num == 0 {
		return ""
	}
	numStr, ok := num2str[num]
	if ok {
		return numStr
	}
	for i := uint64(20); i < num; i += 10 {
		if num-i < 10 {
			return num2str[i] + "-" + num2str[num-i]
		}
		continue
	}
	return ""
}

func say1to999(num uint64) (ret string) {
	hundreds, remainder := num/HUNDRED, num%HUNDRED
	if hundreds > 0 {
		ret += num2str[hundreds] + " " + num2str[HUNDRED]
	}
	remainderStr := say0to99(remainder)
	if ret != "" && remainderStr != "" {
		ret += " "
	}
	ret += remainderStr
	return
}

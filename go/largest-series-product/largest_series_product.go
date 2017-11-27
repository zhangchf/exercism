package lsproduct

import (
	"errors"
	"unicode"
)

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) {
		return -1, errors.New("Invalid input")
	}

	dArr := make([]int, 0)
	for _, v := range digits {
		if !unicode.IsDigit(v) {
			return -1, errors.New("Invalid input")
		}
		dArr = append(dArr, int(v-'0'))
	}

	var largest int64 = 0
	var product int64 = 1
	for i := 0; i <= len(dArr)-span; i++ {
		product = 1
		for j := 0; j < span; j++ {
			product *= int64(dArr[i+j])
		}
		if product > largest {
			largest = product
		}
	}
	return largest, nil
}

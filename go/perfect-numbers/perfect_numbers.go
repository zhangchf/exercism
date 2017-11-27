package perfect

import (
	"errors"
)

const testVersion = 1

var ErrOnlyPositive = errors.New("Only positive input allowed")

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

func Classify(number uint64) (class Classification, err error) {
	if number <= 0 {
		err = ErrOnlyPositive
		return
	}
	var aliquotSum uint64 = 0
	for i := uint64(1); i <= number/2; i++ {
		if number%i == 0 {
			aliquotSum += i
		}
	}
	switch {
	case aliquotSum == number:
		class = ClassificationPerfect
	case aliquotSum > number:
		class = ClassificationAbundant
	case aliquotSum < number:
		class = ClassificationDeficient
	}
	return
}

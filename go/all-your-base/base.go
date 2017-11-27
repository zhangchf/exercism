package allyourbase

import "errors"

var ErrInvalidBase = errors.New("Invalid Base")
var ErrInvalidDigit = errors.New("Invalid Digit")

func ConvertToBase(inputBase uint64, inputDigits []uint64, outputBase uint64) (outputDigits []uint64, err error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, ErrInvalidBase
	}
	inputNum := uint64(0)
	for i, digit := range inputDigits {
		if digit >= inputBase {
			return nil, ErrInvalidDigit
		}
		inputNum += digit
		if i != len(inputDigits)-1 {
			inputNum *= inputBase
		}
	}
	outputDigits = make([]uint64, 0)
	for {
		quotient := inputNum / outputBase
		remainder := inputNum % outputBase
		if quotient == 0 {
			outputDigits = append(outputDigits, remainder)
			break
		} else {
			outputDigits = append(outputDigits, remainder)
			inputNum = quotient
		}
	}
	for i, j := 0, len(outputDigits)-1; i < j; i, j = i+1, j-1 {
		outputDigits[i], outputDigits[j] = outputDigits[j], outputDigits[i]
	}
	return outputDigits, nil
}

package wordy

import (
	"strconv"
	"strings"
)

var ops = []string{"plus", "minus", "divided by", "multiplied by"}

func Answer(question string) (int, bool) {
	question = strings.Trim(question, " ?")
	result := 0
	fields := strings.Fields(question)
	i := 0
	for ; i < len(fields); i++ {
		if num, err := strconv.Atoi(fields[i]); err == nil {
			result = num
			break
		}
	}

	i += 1
	if i >= len(fields) {
		return 0, false
	}

	for i < len(fields) {
		nextOp := fields[i]
		if strings.Compare(nextOp, "plus") == 0 {
			if i+1 >= len(fields) {
				return 0, false
			}
			rightOperator := fields[i+1]
			if num, err := strconv.Atoi(rightOperator); err != nil {
				return 0, false
			} else {
				result += num
			}
			i += 2
		} else if strings.Compare(nextOp, "minus") == 0 {
			if i+1 >= len(fields) {
				return 0, false
			}
			rightOperator := fields[i+1]
			if num, err := strconv.Atoi(rightOperator); err != nil {
				return 0, false
			} else {
				result -= num
			}
			i += 2
		} else if strings.Compare(nextOp, "divided") == 0 {
			if i+2 >= len(fields) {
				return 0, false
			}
			rightOperator := fields[i+2]
			if num, err := strconv.Atoi(rightOperator); err != nil {
				return 0, false
			} else {
				result /= num
			}
			i += 3
		} else if strings.Compare(nextOp, "multiplied") == 0 {
			if i+2 >= len(fields) {
				return 0, false
			}
			rightOperator := fields[i+2]
			if num, err := strconv.Atoi(rightOperator); err != nil {
				return 0, false
			} else {
				result *= num
			}
			i += 3
		} else {
			return 0, false
		}
	}
	return result, true
}

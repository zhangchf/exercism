package change

import (
	"errors"
)

func Change(coins []int, target int) ([]int, error) {
	finalOutput := make([]int, 0)
	if target == 0 {
		return finalOutput, nil
	}
	found := false
	for i := len(coins); i > 0; i-- {
		if target < coins[i-1] {
			continue
		}
		output, innerFound := changeOnSubset(coins[:i], target)
		if innerFound {
			if len(finalOutput) == 0 || len(output) < len(finalOutput) {
				finalOutput = output
			}
			found = true
		}
	}
	if !found {
		return nil, errors.New("Invalid input")
	}
	// swap output
	for i, j := 0, len(finalOutput)-1; i < j; i, j = i+1, j-1 {
		finalOutput[i], finalOutput[j] = finalOutput[j], finalOutput[i]
	}
	return finalOutput, nil
}

func changeOnSubset(coins []int, target int) ([]int, bool) {
	output := make([]int, 0)
	for i := len(coins) - 1; i >= 0 && target > 0; i-- {
		coin := coins[i]
		for target >= coin {
			output = append(output, coin)
			target -= coin
		}
	}

	for target > 0 && len(output) > 0 {
		target += output[0]
		output = output[1:]
		if subOutput, subFound := changeOnSubset(coins[:len(coins)-1], target); subFound {
			output = append(output, subOutput...)
			target = 0
			break
		}
	}

	if target != 0 {
		return nil, false
	}
	return output, true
}

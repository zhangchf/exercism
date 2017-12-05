package collatzconjecture

import "errors"

func CollatzConjecture(input int) (steps int, err error) {
	if input <= 0 {
		return 0, errors.New("Invalid input")
	}
	steps = 0
	for input != 1 {
		if input%2 == 0 {
			input /= 2
		} else {
			input = 3*input + 1
		}
		steps++
	}
	return steps, nil
}

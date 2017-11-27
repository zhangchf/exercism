package grains

import "errors"

const testVersion = 1

func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("Invalid input")
	}

	return 1 << uint(num - 1), nil
}

func Total() uint64 {
	return 1 << uint(64) - 1
}

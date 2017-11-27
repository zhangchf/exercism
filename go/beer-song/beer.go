package beer

import (
	"errors"
	"fmt"
)

func Verse(line int) (string, error) {
	if !isValidInput(line) {
		return "", errors.New("Invalid input")
	}
	if line >= 3 && line <= 99 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", line, line, line-1), nil
	} else if line == 2 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", line, line, line-1), nil
	} else if line == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	} else {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
}

func Verses(upperBound, lowerBound int) (result string, err error) {
	if !isValidInput(upperBound) || !isValidInput(lowerBound) || upperBound < lowerBound {
		return "", errors.New("Invalid input")
	}

	for i := upperBound; i >= lowerBound; i-- {
		if s, err := Verse(i); err == nil {
			if i != upperBound {
				result += "\n"
			}
			result += s
		}
	}
	result += "\n"
	return result, nil
}

func Song() string {
	s, _ := Verses(99, 0)
	return s
}

func isValidInput(line int) bool {
	return line >= 0 && line <= 99
}

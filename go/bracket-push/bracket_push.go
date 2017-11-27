// Package brackets check for the validity of parenthesis inside a input string
package brackets

import "fmt"

const testVersion = 5

var counterp = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

// Bracket check the validity of brackets of the given string
func Bracket(input string) (valid bool, err error) {
	fmt.Println(input)
	stack := make([]int32, 0)
	for _, c := range input {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if peek(stack) != counterp[c] {
				return false, nil
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false, nil
	}
	return true, nil
}

func peek(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	return arr[len(arr)-1]
}

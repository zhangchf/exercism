// Package anagram have anagram related functions
package anagram

import (
	"fmt"
	"strings"
)

const testVersion = 2

// Detect find anagrams of a given subject from the given candidates
func Detect(subject string, candidates []string) (anagrams []string) {
	fmt.Println(subject)
	subject = strings.ToLower(subject)
	letters := getLetters(subject)
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if len(lowerCandidate) != len(subject) || subject == lowerCandidate {
			continue
		}
		if isAnagram(lowerCandidate, letters) {
			anagrams = append(anagrams, candidate)
		}
	}
	return
}

// getLetters get the occurences of each letter in the given input string
func getLetters(input string) map[rune]int {
	var result = make(map[rune]int)
	for _, r := range input {
		result[r] += 1
	}
	return result
}

// isAnagram checks whether the given subject is an anagram of the given letter occurence map.
func isAnagram(subject string, letters map[rune]int) bool {
	subject = strings.ToLower(subject)
	for k, v := range letters {
		for i, pos := 0, 0; i < v; i++ {
			if ind := strings.IndexRune(subject[pos:], k); ind != -1 {
				pos = ind + 1
			} else {
				return false
			}
		}
	}
	return true
}

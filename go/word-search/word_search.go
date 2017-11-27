package wordsearch

import (
	"errors"
	"strings"
)

var errInvalidInput = errors.New("Invalid Input")

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	puzzleRows, puzzleCols := len(puzzle), len(puzzle[0])
	for _, row := range puzzle {
		if len(row) != puzzleCols {
			return nil, errInvalidInput
		}
	}
	
	result := map[string][2][2]int{}
	found := false
	for _, word := range words {
		reversedWord := reverse(word)
		for r, row := range puzzle {
			if ind := strings.Index(row, word); ind != -1 {
				result[word] = [2][2]int{{ind, r}, {ind+len(word)-1, r}} 
				found = true
				continue
			}
			if ind := strings.Index(row, reversedWord); ind != -1 {
				result[word] = [2][2]int{{ind+len(word)-1, r},{ind, r}}
				found = true
				continue
			}
		}
		
		for i := 0; i < puzzleCols; i++ {
			col := make([]byte, puzzleRows)
			for j := 0; j < puzzleRows; j++ {
				col[j] = puzzle[j][i]
			}
			colStr := string(col)
			if ind := strings.Index(colStr, word); ind != -1 {
				result[word] = [2][2]int{{i, ind}, {i, ind+len(word)-1}}
				found = true
				continue
			}
			if ind := strings.Index(colStr, reversedWord); ind != -1 {
				result[word] = [2][2]int{{i, ind+len(word)-1}, {i, ind}}
				found = true
				continue
			}
		}
		
		for i := 0; i < puzzleRows; i++ {
			col := make([]byte, 0)
			for r, c := i, 0; r < puzzleRows && c < puzzleCols; r, c = r+1, c+1 {
				col = append(col, puzzle[r][c])
			}
			colStr := string(col)
			if ind := strings.Index(colStr, word); ind != -1 {
				result[word] = [2][2]int{{i+ind, i+ind}, {i+ind+len(word)-1, i+ind+len(word)-1}}
				found = true
				continue
			}
			if ind := strings.Index(colStr, reversedWord); ind != -1 {
				result[word] = [2][2]int{{i+ind+len(word)-1, i+ind+len(word)-1}, {i+ind, i+ind}}
				found = true
				continue
			}
		}

		for i := 0; i < puzzleCols; i++ {
			col := make([]byte, 0)
			for r, c := 0, i; r < puzzleRows && c < puzzleCols; r, c = r+1, c+1 {
				col = append(col, puzzle[r][c])
			}
			colStr := string(col)
			if ind := strings.Index(colStr, word); ind != -1 {
				result[word] = [2][2]int{{i+ind, i+ind}, {i+ind+len(word)-1, i+ind+len(word)-1}}
				found = true
				continue
			}
			if ind := strings.Index(colStr, reversedWord); ind != -1 {
				result[word] = [2][2]int{{i+ind+len(word)-1, i+ind+len(word)-1}, {i+ind, i+ind}}
				found = true
				continue
			}
		}
	}
	if !found {
		return nil, errInvalidInput
	}
	return result, nil	
}

func reverse(input string) string {
	lenInput := len(input)
	out := make([]byte, len(input))
	for i, b := range input {
		out[lenInput-1-i] = byte(b)
	}
	return string(out)
}

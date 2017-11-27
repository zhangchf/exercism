package minesweeper

import (
	"bytes"
	"errors"
)

type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

func (b Board) Count() error {
	if !b.checkBorder() {
		return errors.New("Invalid board")
	}
	rows, cols := len(b), len(b[0])
	for r := 1; r <= rows-2; r++ {
		for c := 1; c <= cols-2; c++ {
			if b[r][c] == '*' {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						tr, tc := r+i, c+j
						if tr >= 1 && tr <= rows-2 && tc >= 1 && tc <= cols-2 && b[tr][tc] != '*' {
							if b[tr][tc] == ' ' {
								b[tr][tc] = '0'
							}
							b[tr][tc]++
						}
					}
				}
			}
		}
	}
	return nil
}

func (b Board) checkBorder() bool {
	rows, cols := len(b), len(b[0])
	for row, line := range b {
		if len(line) != cols {
			return false
		}
		if row == 0 || row == rows-1 {
			for i, c := range line {
				if i == 0 || i == cols-1 {
					if c != '+' {
						return false
					}
				} else if c != '-' {
					return false
				}
			}
		} else {
			for i, c := range line {
				if i == 0 || i == cols-1 {
					if c != '|' {
						return false
					}
				} else if c != '*' && c != ' ' {
					return false
				}
			}
		}
	}
	return true
}

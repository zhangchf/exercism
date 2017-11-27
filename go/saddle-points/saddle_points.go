package matrix

import (
	"strings"
	"strconv"
	"errors"
)

type Pair struct {
	row, col int
}

type Matrix [][]int


func New(mStr string) (*Matrix, error) {
	mStr = strings.TrimSpace(mStr)
	rows := strings.Split(mStr, "\n")
	matrix := Matrix{}
	for _, row := range rows {
		cols := strings.Split(row, " ")
		rowArray := make([]int, len(cols))
		for c, col := range cols {
			if num, err := strconv.Atoi(col); err == nil {
				rowArray[c] = num
			} else {
				return nil, errors.New("Invalid row:" + col)
			}
		}
		matrix = append(matrix, rowArray)
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	result := []Pair{}
	for r, row := range *m {
		maxInRow := row[0]
		saddleCol := 0
		for j := 0; j < len(row); j++ {
			if row[j] > maxInRow {
				maxInRow = row[j]
				saddleCol = j
			}
		}
		isSaddle := true
		for i := 0; i < len(*m); i++ {
			if (*m)[i][saddleCol] < maxInRow {
				isSaddle = false
				break
			}
		}
		if isSaddle {
			result = append(result, Pair{r, saddleCol})
		}
	}
	return result
}
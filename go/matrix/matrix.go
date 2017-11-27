package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	data       []int
	rows, cols int
}

func New(s string) (m *Matrix, err error) {
	m = &Matrix{}
	rowsStr := strings.Split(s, "\n")
	m.rows = len(rowsStr)
	for _, rowStr := range rowsStr {
		rowStr = strings.TrimSpace(rowStr)
		colParts := strings.Split(rowStr, " ")
		clen := len(colParts)
		if m.cols == 0 {
			m.cols = clen
		} else if m.cols != clen {
			return nil, errors.New("Invalid input")
		}
		col := make([]int, m.cols)
		for i := 0; i < m.cols; i++ {
			v, err := strconv.Atoi(colParts[i])
			if err != nil {
				return nil, err
			}
			col[i] = v
		}
		m.data = append(m.data, col...)
	}
	return m, nil
}

func (m *Matrix) Rows() (rows [][]int) {
	for i := 0; i < m.rows; i++ {
		row := make([]int, 0)
		row = append(row, m.data[i*m.cols:i*m.cols+m.cols]...)
		rows = append(rows, row)
	}
	return rows
}

func (m *Matrix) Cols() (cols [][]int) {
	for i := 0; i < m.cols; i++ {
		col := make([]int, m.rows)
		for j := 0; j < m.rows; j++ {
			col[j] = m.data[j*m.cols+i]
		}
		cols = append(cols, col)
	}
	return cols
}

func (m *Matrix) Set(r, c, v int) bool {
	if r < 0 || r >= m.rows || c < 0 || c >= m.cols {
		return false
	}
	m.data[r*m.cols+c] = v
	return true
}

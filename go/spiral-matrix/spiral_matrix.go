package spiralmatrix

import "math"

type Matrix struct {
	data [][]int
}

func SpiralMatrix(size int) [][]int {
	m := &Matrix{}
	m.data = make([][]int, size)
	for i := 0; i < size; i++ {
		m.data[i] = make([]int, size)
	}
	nextInt := 1
	for i := 0; i < int(math.Ceil(float64(size)/2)); i++ {
		nextInt = fillBorder(m, nextInt, i, size-1-i, i, size-1-i)
	}
	return m.data
}

func fillBorder(m *Matrix, startNum int, l, r, t, b int) (nextInt int) {
	for i := l; i <= r; i++ {
		m.data[t][i] = startNum
		startNum++
	}
	for j := t + 1; j <= b; j++ {
		m.data[j][r] = startNum
		startNum++
	}
	for i := r - 1; i >= l; i-- {
		m.data[b][i] = startNum
		startNum++
	}
	for j := b - 1; j >= t+1; j-- {
		m.data[j][l] = startNum
		startNum++
	}
	return startNum
}

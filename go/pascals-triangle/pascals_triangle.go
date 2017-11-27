package pascal

const testVersion = 1

func Triangle(size int) [][]int {
	var result = make([][]int, size)
	if size == 1 {
		result[size-1] = []int{1}
	} else {
		result = Triangle(size - 1)
		var newLine = make([]int, size)
		for i := 0; i < size; i++ {
			preLine := result[size-2]
			if i == 0 {
				newLine[i] = preLine[i]
			} else if i == size-1 {
				newLine[i] = preLine[i-1]
			} else {

				newLine[i] = preLine[i-1] + preLine[i]
			}
		}
		result = append(result, newLine)
	}

	return result
}

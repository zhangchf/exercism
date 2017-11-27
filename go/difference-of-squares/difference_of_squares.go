package diffsquares

const testVersion = 1

func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}

func SquareOfSums(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

func Difference(num int) int {
	return SquareOfSums(num) - SumOfSquares(num)
}

package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors... int) (result int) {
	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if i % div == 0 {
				result += i
				break
			}
		}
	}
	return
}

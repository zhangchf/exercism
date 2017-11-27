package prime

var primes = []int {2}

func Nth(nth int) (int, bool) {
	if nth < 1 {
		return 0, false
	}
	for len(primes) < nth {
		last := primes[len(primes)-1]
		for i := last+1; ;i++ {
			isPrime := true
			for _, v := range primes {
				if i%v == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primes = append(primes, i)
				break
			}
		}
	}
	return primes[nth-1], true
}

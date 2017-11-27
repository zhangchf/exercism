package sieve

const testVersion = 1

var compMap = make(map[int]bool)

var primes = []int {2}

func Sieve(limit int) []int {
	for pi := 0; pi < len(primes); pi++ {
		p := primes[pi]
		for i := 2; i*p <= limit; i++ {
			compMap[i*p] = true
		}
		// Find next prime
		curMaxPrime := primes[len(primes)-1]
		for j := curMaxPrime+1; j <= limit; j++ {
			if compMap[j] {
				continue
			} else {
				primes = append(primes, j)
				break
			}
		}
	}
	return primes
}

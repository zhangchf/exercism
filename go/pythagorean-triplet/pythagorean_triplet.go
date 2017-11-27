package pythagorean

const testVersion = 1

type Triplet []int

func Range(min, max int) (ts []Triplet) {
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i + j*j == k*k {
					ts = append(ts, Triplet{i, j, k})
				}
			}
		}
	}
	return
}

func Sum(sum int) (ts []Triplet) {
	for i := 1; i <= sum; i++ {
		for j := i; j <= sum-i; j++ {
			k := sum-i-j
			if i*i + j*j == k*k {
				ts = append(ts, Triplet{i, j, k})
			}
		}
	}
	return
}

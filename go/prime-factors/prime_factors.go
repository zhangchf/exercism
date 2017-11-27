package prime

func Factors(in int64)(out []int64) {
	out = make([]int64, 0)
	rem := in
	for i := int64(2); i <= rem;  {
		if rem%i == 0 {
			out = append(out, i)
			rem = rem/i
		} else {
			i++
		}
	}
	return
}

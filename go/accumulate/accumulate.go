package accumulate

const testVersion = 1

func Accumulate(original []string, op func(string) string) []string {
	result := make([]string, len(original))
	for i := range original {
		result[i] = op(original[i])
	}
	return result
}

package series

const testVersion = 2

func All(size int, s string) []string {
	var result []string
	if len(s) >= size {
		for i := 0; i <= len(s)-size; i++ {
			result = append(result, s[i:i+size])
		}
	}
	return result
}

func UnsafeFirst(size int, s string) string {
	if len(s) < size {
		return s
	} else {
		return s[0:size]
	}
}

func First(size int, s string) (first string, ok bool) {
	if len(s) < size || size <= 0 {
		return "", false
	} else {
		return s[0:size], true
	}
}

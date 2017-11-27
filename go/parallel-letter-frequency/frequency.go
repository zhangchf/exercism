package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(inputs []string) FreqMap {
	resCh := make(chan FreqMap)
	for _, s := range inputs {
		go func(str string) {
			resCh <- Frequency(str)
		}(s)
	}

	result := make(FreqMap)
	for range inputs {
		for k, v := range <-resCh {
			result[k] += v
		}
	}
	return result
}

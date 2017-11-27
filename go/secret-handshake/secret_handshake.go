package secret

const testVersion = 2

var handshakes = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(input uint) []string {
	var result []string

	for i, v := range handshakes {
		if input&(1<<uint(i)) != 0 {
			result = append(result, v)
		}
	}

	if input&(1<<uint(len(handshakes))) != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}

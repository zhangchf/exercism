package cryptosquare

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func EncodeOld(plain string) (cipher string) {
	// Step 1: Normalize and downcase
	plain = strings.ToLower(plain)
	plain = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, plain)

	fmt.Println("Normalized string:", plain)

	// Step 2: Break into rows
	col, row := 0, 0
	plen := len(plain)
	for i := 0; i <= plen; i++ {
		if i*i >= plen {
			col, row = i, i
			break
		} else if i*i < plen && i*(i+1) >= plen {
			col, row = i+1, i
			break
		}
	}
	fmt.Println("col, row:", col, row)
	splits := make([]string, row)
	for i, j := 0, 0; i < plen; i, j = i+col, j+1 {
		end := i + col
		if end > plen {
			end = plen
		}
		splits[j] = plain[i:end]
	}

	fmt.Println("Splited rectange:", splits)

	// Transforming
	transform := make([]string, col)
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if i < len(splits[j]) {
				transform[i] += splits[j][i : i+1]
			}
		}
	}

	fmt.Printf("Transformed rectange: %v\n", transform)
	fmt.Println()

	return strings.Join(transform, " ")
}

func Encode(plain string) (cipher string) {
	var inBuf bytes.Buffer

	for _, v := range plain {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			inBuf.WriteRune(unicode.ToLower(v))
		}
	}

	col := int(math.Ceil(math.Sqrt(float64(inBuf.Len()))))

	var outBuf bytes.Buffer
	inBytes := inBuf.Bytes()
	for i := 0; i < col; i++ {
		for j := i; j < len(inBytes); j += col {
			outBuf.WriteByte(inBytes[j])
		}
		if col > i+1 {
			outBuf.WriteByte(' ')
		}
	}

	return outBuf.String()
}

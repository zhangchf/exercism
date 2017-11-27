package palindrome

import (
	"errors"
	"strconv"
)

const testVersion = 1

type Product struct {
	product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			pv := i * j
			if !isPalindrome(pv) {
				continue
			}
			if pv == pmin.product {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}
			if pv < pmin.product || pmin.product == 0 {
				pmin.product = pv
				pmin.Factorizations = pmin.Factorizations[:0]
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}

			if pv == pmax.product {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}
			if pv > pmax.product || pmax.product == 0 {
				pmax.product = pv
				pmax.Factorizations = pmax.Factorizations[:0]
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}
		}
	}
	if pmin.product == 0 && pmax.product == 0 {
		err = errors.New("no palindromes")
	}
	return
}

func isPalindrome(value int) bool {
	sv := strconv.Itoa(value)
	for i, j := 0, len(sv)-1; i < j; i, j = i+1, j-1 {
		if sv[i] != sv[j] {
			return false
		}
	}
	return true
}

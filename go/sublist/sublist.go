package sublist

import (
	"strings"
	"strconv"
)

type Relation string

const (
	EQUAL Relation = "equal"
	UNEQUAL = "unequal"
	SUBLIST = "sublist"
	SUPERLIST = "superlist"
)

func Sublist(first, second []int) Relation {
	s1, s2 := "", ""
	for _, v := range first {
		s1 += strconv.Itoa(v) + " "
	}
	for _, v := range second {
		s2 += strconv.Itoa(v) + " "
	}
	if strings.Compare(s1, s2) == 0 {
		return EQUAL
	}
	if strings.Contains(s1, s2) {
		return SUPERLIST
	}
	if strings.Contains(s2, s1) {
		return SUBLIST
	}
	return UNEQUAL
}


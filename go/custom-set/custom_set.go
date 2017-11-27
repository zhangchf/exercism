package stringset

import (
	"fmt"
	"strings"
)

const testVersion = 4

type Set struct {
	data []string
}

type SetA []string

func New() Set {
	return Set{}
}

func NewFromSlice(slice []string) Set {
	set := Set{}
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

func (s Set) String() (res string) {
	res += "{"
	for i := 0; i < len(s.data); i++ {
		if i != 0 {
			res += ", "
		}
		res += fmt.Sprintf("%q", s.data[i])
	}
	res += "}"
	return
}

func (s *Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set) Has(elem string) bool {
	for _, v := range s.data {
		if strings.Compare(v, elem) == 0 {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		s.data = append(s.data, elem)
	}
}

func Subset(setA, setB Set) bool {
	for _, v := range setA.data {
		if !setB.Has(v) {
			return false
		}
	}
	return true
}

func Disjoint(setA, setB Set) bool {
	for _, v := range setA.data {
		if setB.Has(v) {
			return false
		}
	}
	return true
}

func Equal(setA, setB Set) bool {
	if len(setA.data) != len(setB.data) {
		return false
	}
	return Subset(setA, setB)
}

func Intersection(setA, setB Set) Set {
	interSet := Set{}
	for _, v := range setA.data {
		if setB.Has(v) {
			interSet.data = append(interSet.data, v)
		}
	}
	return interSet
}

func Difference(setA, setB Set) Set {
	diffSet := Set{}
	for _, v := range setA.data {
		if !setB.Has(v) {
			diffSet.data = append(diffSet.data, v)
		}
	}
	return diffSet
}

func Union(setA, setB Set) Set {
	unionSet := Set{}
	appendSet := &setA
	addSet := &setB
	if len(setB.data) > len(setA.data) {
		appendSet, addSet = addSet, appendSet
	}
	unionSet.data = append(unionSet.data, appendSet.data...)
	for _, v := range addSet.data {
		unionSet.Add(v)
	}
	return unionSet
}

package strain

type Ints []int
type Strings []string
type Lists []Ints

func (ints Ints) Keep(f func(int) bool) (out Ints) {
	for _, v := range ints {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

func (ints Ints) Discard(f func(int) bool) (out Ints) {
	for _, v := range ints {
		if !f(v) {
			out = append(out, v)
		}
	}
	return
}

func (ss Strings) Keep(f func(string) bool) (out Strings) {
	for _, v := range ss {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

func (ss Strings) Discard(f func(string) bool) (out Strings) {
	for _, v := range ss {
		if !f(v) {
			out = append(out, v)
		}
	}
	return
}

func (ls Lists) Keep(f func([]int) bool) (out Lists) {
	for _, v := range ls {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

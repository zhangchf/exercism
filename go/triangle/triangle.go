package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if !(isValidSide(a) && isValidSide(b) && isValidSide(c)) {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

func isValidSide(v float64) bool {
	return v > 0 && v != math.Inf(1)
}

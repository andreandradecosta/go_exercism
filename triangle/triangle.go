//+build !example

package triangle

import "math"

// Kind represents a type of a triangle
type Kind string

// Types of triangles
const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func validNumbers(a, b, c float64) bool {
	return !math.IsNaN(a) &&
		!math.IsNaN(b) &&
		!math.IsNaN(c) &&
		!math.IsInf(a, 1) &&
		!math.IsInf(b, 1) &&
		!math.IsInf(c, 1) &&
		!math.IsInf(a, -1) &&
		!math.IsInf(b, -1) &&
		!math.IsInf(c, -1)
}

// KindFromSides returns the type of a triangle based on its sides
func KindFromSides(a, b, c float64) Kind {
	if !validNumbers(a, b, c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

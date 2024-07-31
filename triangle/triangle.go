// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind rune

const (
	// Pick values for the following identifiers used by the test program.
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a == 0 || b == 0 || c == 0 || !(a+b >= c) || !(b+c >= a) || !(a+c >= b) {
		k = NaT
	} else {
		switch {
		case a == b && a == c:
			k = Equ
		case a == b || a == c || b == c:
			k = Iso
		case a != b && a != c && b != c:
			k = Sca
		default:
			k = NaT
		}
	}
	return k
}
// Triangle task of Exercism.io Go track.
package triangle

import "math"

// Type of triangle.
type Kind int

// Different types of triangles.
const (
	NaT Kind = iota
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
)

// Stringify triangle type.
func (k Kind) String() string {
	return []string{"NaT", "Equ", "Iso", "Sca"}[k]
}

// Determine kind of a triangle.
func KindFromSides(a, b, c float64) Kind {
	// order does matter
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT
	case a <= 0.0 || b <= 0.0 || c <= 0.0:
		return NaT
	case a+b < c || b+c < a || c+a < b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	}
	return Sca
}

//// Another take on determining triangle kind.
//// The "func(a, b, c float64) bool" looks too verbose.
//func k2(a, b, c float64) Kind {
//	x := [...]struct {
//		k Kind
//		t func(a, b, c float64) bool
//	}{ // order does matter
//		{NaT, func(a, b, c float64) bool { return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)          }},
//		{NaT, func(a, b, c float64) bool { return math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) }},
//		{NaT, func(a, b, c float64) bool { return a <= 0.0 || b <= 0.0 || c <= 0.0                         }},
//		{NaT, func(a, b, c float64) bool { return a+b < c || b+c < a || c+a < b                            }},
//		{Equ, func(a, b, c float64) bool { return a == b && b == c                                         }},
//		{Iso, func(a, b, c float64) bool { return a == b || b == c || c == a                               }}}
//
//	for _, t := range x {
//		if t.t(a, b, c) {
//			return t.k
//		}
//	}
//
//	return Sca
//}

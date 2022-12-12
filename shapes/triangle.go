package shapes

import "math"

type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	// find area of triangle using side lengths
	// Heron's formula
	s := (t.Perimeter()) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

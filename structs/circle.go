package structs

import "math"

// Circle represents a circle by its radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of a Circle struct.
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

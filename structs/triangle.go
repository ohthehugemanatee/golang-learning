package structs

// Circle represents a circle by its radius.
type Triangle struct {
	Height float64
	Width float64
}

// Area calculates the area of a Circle struct.
func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}

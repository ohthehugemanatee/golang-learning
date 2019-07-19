package structs

// Rectangle defines... a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a square given height and width.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

package rectangle

// Perimeter calculates the perimeter of a square given height and width.
func Perimeter(height, width float64) float64 {
	return 2 * (width + height)
}

func Area(height, width float64) float64 {
	return height * width
}

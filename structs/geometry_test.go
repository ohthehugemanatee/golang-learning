package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 6.3}
	result := rectangle.Perimeter()
	expected := 22.6

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %.2f, expected %.2f", got, want)
		}
	}
	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.3}
		checkArea(t, rectangle, 31.5)
	})
	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

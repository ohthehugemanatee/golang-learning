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
	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.3}
		result := rectangle.Area()
		expected := 31.5

		if result != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, result)
		}
	})
	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expected := 314.1592653589793

		if result != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, result)
		}
	})
}

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
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 5.0, Height: 6.3}, hasArea: 31.5},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Width: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v Got %.2f, expected %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

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
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 5.0, Height: 6.3}, want: 31.5},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Width: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Got %.2f, expected %.2f", got, tt.want)
		}
	}
}

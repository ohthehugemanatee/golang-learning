package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 6.3}
	result := Perimeter(rectangle)
	expected := 22.6

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.0, 6.3}
	result := Area(rectangle)
	expected := 31.5

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

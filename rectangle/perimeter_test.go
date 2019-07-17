package rectangle

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	result := Perimeter(5.0, 6.3)
	expected := 22.6

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	result := Area(5.0, 6.3)
	expected := 31.5

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

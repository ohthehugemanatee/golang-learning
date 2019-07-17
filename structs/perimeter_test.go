package perimeter

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

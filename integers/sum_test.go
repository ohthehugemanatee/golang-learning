package integers

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(3, 4)
	expected := 7
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

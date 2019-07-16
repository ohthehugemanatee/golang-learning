package integers

import "testing"

func TestSum(t *testing.T) {
	t.Run("slice of 3 numbers", func(t *testing.T) {
		numbers := []int{5, 6, 7}
		sum := Sum(numbers)
		expected := 18
		if sum != expected {
			t.Errorf("expected '%d' but got '%d', gave %v", expected, sum, numbers)
		}
	})
}

package integers

import "testing"

func TestSum(t *testing.T) {
	t.Run("slice of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("slice of 3 numbers", func(t *testing.T) {
		numbers := []int{5, 6, 7}
		sum := Sum(numbers)
		expected := 18
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

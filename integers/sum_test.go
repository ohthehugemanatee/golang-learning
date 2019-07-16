package integers

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7})
	expected := []int{6, 22}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSumAllTails(t *testing.T) {
	result := SumAllTails([]int{1, 2, 3, 4}, []int{5, 6}, []int{7, 8, 9})
	expected := []int{9, 6, 17}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

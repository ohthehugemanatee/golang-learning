package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"
	if actual != expected {
		t.Errorf("Expected %q but received %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("f", 3)
	fmt.Println(result)
	// Output: fff
}

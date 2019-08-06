package reflection

import "testing"

func TestWalk(t *testing.T) {
	arr := struct {
		string
	}{
		"value1",
	}
	got := []string{}
	Walk(arr, func(str string) {
		got = append(got, str)
	})
	if len(got) != 1 {
		t.Errorf("Function called the wrong number of times. Expected %d, got %d", 1, len(got))
	}
}

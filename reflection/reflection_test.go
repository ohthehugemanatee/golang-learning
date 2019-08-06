package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "value1"
	arr := struct {
		string
	}{
		expected,
	}
	got := []string{}
	Walk(arr, func(str string) {
		got = append(got, str)
	})
	if len(got) != 1 {
		t.Errorf("Function called the wrong number of times. Expected %d, got %d", 1, len(got))
	}
	if got[0] != expected {
		t.Errorf("Wrong value in function call. Expected %s, got %s", expected, got[0])
	}
}

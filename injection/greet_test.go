package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Cam")
	got := buffer.String()
	expected := "Hello, Cam"
	if got != expected {
		t.Errorf("Got %q but expected %q", got, expected)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Cam")
	want := "Hello, Cam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

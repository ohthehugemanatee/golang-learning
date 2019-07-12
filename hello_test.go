package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Cam", "")
		want := "Hello, Cam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("and now in Spanish", func(t *testing.T) {
		got := Hello("Cam", "Spanish")
		want := "Hola, Cam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("and now in French", func(t *testing.T) {
		got := Hello("Cam", "French")
		want := "Bonjour, Cam"
		assertCorrectMessage(t, got, want)
	})

}

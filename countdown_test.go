package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
	sleeperCalls := spySleeper.Calls
	sleeperCallsWanted := 3
	if sleeperCalls != sleeperCallsWanted {
		t.Errorf("Sleeper was called %q times. Expected %q times.", sleeperCalls, sleeperCallsWanted)
	}
}

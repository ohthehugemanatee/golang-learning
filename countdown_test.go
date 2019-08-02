package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type OperationsSpy struct {
	Calls []string
}

func (s *OperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *OperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("Output is correct", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &OperationsSpy{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
		sleeperCalls := len(spySleeper.Calls)
		sleeperCallsWanted := 3
		if sleeperCalls != sleeperCallsWanted {
			t.Errorf("Sleeper was called %q times. Expected %q times.", sleeperCalls, sleeperCallsWanted)
		}
	})
	t.Run("Wait is called between each output", func(t *testing.T) {
		operationsSpy := &OperationsSpy{}
		Countdown(operationsSpy, operationsSpy)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, operationsSpy.Calls) {
			t.Errorf("Sleep and write were not called in alternating order. Call order was %v", operationsSpy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("Expected sleep time of %q, got %q", sleepTime, spyTime.durationSlept)
	}
}

const (
	write = "write"
	sleep = "sleep"
)

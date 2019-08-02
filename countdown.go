package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is anything with a sleep function.
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown from 3
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

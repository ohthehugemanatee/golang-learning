package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown from 3
func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}

package injection

import (
	"fmt"
	"io"
)

// Greet the user by name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

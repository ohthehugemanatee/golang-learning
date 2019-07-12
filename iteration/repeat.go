package iteration

// Repeat just repeats the argument 5 times.
func Repeat(char string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += char
	}
	return
}

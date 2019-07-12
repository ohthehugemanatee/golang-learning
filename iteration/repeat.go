package iteration

// Repeat just repeats the argument 5 times.
func Repeat(char string, iterations int) (repeated string) {
	for i := 0; i < iterations; i++ {
		repeated += char
	}
	return
}

package integers

// Sum adds an array of 5 numbers together.
func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

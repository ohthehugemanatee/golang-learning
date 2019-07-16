package integers

// Sum adds an array of 5 numbers together.
func Sum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}

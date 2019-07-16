package integers

// Sum adds an array of 5 numbers together.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll returns sums of all sets you pass in.
func SumAll(sets ...[]int) (sums []int) {
	for _, set := range sets {
		sums = append(sums, Sum(set))
	}
	return
}

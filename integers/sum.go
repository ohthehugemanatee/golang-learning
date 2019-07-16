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

// SumAllTails returns the sums of all values except the HEAD of each array.
func SumAllTails(sets ...[]int) (sums []int) {
	for _, set := range sets {
		sum := Sum(set[1:])
		sums = append(sums, sum)
	}
	return
}

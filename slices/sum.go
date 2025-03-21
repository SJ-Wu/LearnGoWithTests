package slices

// Sum calculates and returns the total sum of integers in the given slice.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes multiple slices of integers and returns a slice containing the sums of each individual slice.
func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

// SumAllTails calculates the sums of the "tails" of each slice in the variadic input and returns them as a new slice.
// The tail of a slice is defined as all elements except the first.
// An empty slice results in a sum of zero.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

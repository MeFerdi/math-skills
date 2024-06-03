package statmath

import (
	"math"
)

// Median calculates the median of a given slice of float64 values.
func Median(numbers []float64) float64 {
	// Get the length of the numbers slice.
	n := len(numbers)

	// Implement bubble sort to sort the numbers slice.
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If the current number is greater than the next number, swap them.
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// If the length of the numbers slice is even, calculate the median as the average of the two middle numbers.
	if n%2 == 0 {
		// Calculate the median as the average of the two middle numbers.
		median := (numbers[n/2-1] + numbers[n/2]) / 2
		// If the median is halfway between two integers, round up to the nearest integer.
		if median-math.Floor(median) >= 0.5 {
			return math.Ceil(median)
		} else {
			// Otherwise, round down to the nearest integer.
			return math.Floor(median)
		}
	} else {
		// If the length of the numbers slice is odd, the median is the middle number.
		// Convert the float64 to an integer and return it.
		return float64(int(numbers[n/2]))
	}
}

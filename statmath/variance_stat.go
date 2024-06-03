package statmath

import (
	"math"
)

// Variance calculates the variance of a given slice of float64 values.
func Variance(numbers []float64) float64 {
	// Initialize a variable to store the sum of all numbers.
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	// Calculate the mean by dividing the sum by the number of values.
	mean := float64(sum) / float64(len(numbers))

	// Initialize a variable to store the sum of squared differences from the mean.
	sum = 0

	// Calculate the sum of squared differences from the mean.
	for _, num := range numbers {

		diff := float64(num - mean)

		sum += math.Pow(diff, 2)
	}

	return float64(sum) / float64(len(numbers))
}

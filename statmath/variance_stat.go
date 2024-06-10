package statmath

import (
	"math"
)

// Variance calculates the variance of a given slice of float64 values.
func Variance(numbers []float64) float64 {
	// Calculate the mean.
	mean := Average(numbers)

	// Initialize a variable to store the sum of squared differences from the mean.
	variance := 0.0

	// Calculate the sum of squared differences from the mean.
	for _, num := range numbers {
		diff := float64(num - mean)
		variance += math.Pow(diff, 2)
	}

	return float64(variance) / float64(len(numbers))
}

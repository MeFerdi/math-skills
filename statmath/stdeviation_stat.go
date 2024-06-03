package statmath

import (
	"math"
)

// StandardDeviation calculates the standard deviation of a given slice of float64 values.
func StandardDeviation(data []float64) float64 {
	// Get the length of the data slice.
	n := float64(len(data))
	// If the length of the data slice is 0, return 0.
	if n == 0 {
		return 0
	}

	// Calculate the mean of the data.
	var mean float64
	for _, value := range data {
		mean += value
	}
	mean /= n

	// Calculate the variance of the data.
	var variance float64
	for _, value := range data {
		variance += (value - mean) * (value - mean)
	}
	variance /= n

	// Calculate the square root of the variance to get the standard deviation.
	return math.Sqrt(variance)
}

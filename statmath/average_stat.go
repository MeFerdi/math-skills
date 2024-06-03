package statmath

import (
	"math"
)

// Average calculates the average of a given slice of float64 values.
func Average(data []float64) float64 {
	var sum float64

	// Iterate over each value in the data slice.
	for _, value := range data {
		sum += value
	}

	// Calculate the average by dividing the sum by the number of values.
	avg := float64(sum) / float64(len(data))

	// If the average is halfway between two integers (i.e., 0.5 or greater), round up to the nearest integer.
	if math.Mod(avg, 1) >= 0.5 {
		//
		return float64(math.Ceil(avg))
	}

	return float64(math.Floor(avg))
}

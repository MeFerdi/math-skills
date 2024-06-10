package statmath

import (
	"math"
)

// StandardDeviation calculates the standard deviation of a given slice of float64 values.
func StandardDeviation(data []float64) float64 {
	// Calculate the square root of the variance to get the standard deviation.
	return math.Sqrt(Variance(data))
}

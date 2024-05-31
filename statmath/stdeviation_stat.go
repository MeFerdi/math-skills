package statmath

import (
	"math"
)

func StandardDeviation(data []float64) float64 {
	n := float64(len(data))
	if n == 0 {
		return 0
	}

	var mean float64
	for _, value := range data {
		mean += value
	}
	mean /= n

	var variance float64
	for _, value := range data {
		variance += (value - mean) * (value - mean)
	}
	variance /= n

	return math.Sqrt(variance)
}

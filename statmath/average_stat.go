package statmath

import (
	"math"
)

func Average(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	avg := float64(sum) / float64(len(data))

	if math.Mod(avg, 1) >= 0.5 {
		return float64(math.Ceil(avg))
	}
	return float64(math.Floor(avg))
}

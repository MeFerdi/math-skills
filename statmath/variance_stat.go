package statmath

import (
	"math"
)

func Variance(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	mean := float64(sum) / float64(len(numbers))
	sum = 0
	for _, num := range numbers {
		sum += math.Pow(float64(num-mean), 2)
	}
	return float64(sum) / float64(len(numbers))
}

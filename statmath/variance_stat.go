package statmath

import (
	"math"
)

func Variance(numbers []float64) float64 {
	avg := Average(numbers)
	sum := 0
	for _, num := range numbers {
		sum += int(math.Pow(float64(num-avg), 2))
	}
	return float64(sum) / float64(len(numbers))
}

func AverageVal(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

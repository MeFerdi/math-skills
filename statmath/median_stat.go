package statmath

import (
	"math"
)

func Median(numbers []float64) float64 {
	n := len(numbers)

	// Implement bubble sort to sort the numbers slice
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	if n%2 == 0 {
		median := (numbers[n/2-1] + numbers[n/2]) / 2
		if median-math.Floor(median) >= 0.5 {
			return math.Ceil(median)
		} else {
			return math.Floor(median)
		}
	} else {
		return float64(int(numbers[n/2]))
	}
}

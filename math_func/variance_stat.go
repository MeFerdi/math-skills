package main

import (
	"fmt"
	"math"
)

func Variance(numbers []float64) float64 {
	avg := Average(numbers)
	var sum float64
	for _, num := range numbers {
		sum += (num - avg) * (num - avg)
	}
	return sum / float64(len(numbers))
}

func Average(numbers []float64) float64 {
	return math.Sum(numbers) / float64(len(numbers))
}

func main() {
	numbers := []float64{3.2, 5.8, 2.1, 9.7, 4.5}
	variance := Variance(numbers)
	fmt.Printf("Variance: %.2f\n", variance)
}

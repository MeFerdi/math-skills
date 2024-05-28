package main

import (
	"fmt"
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
		return (numbers[n/2-1] + numbers[n/2]) / 2
	} else {
		return numbers[n/2]
	}
}

func main() {
	numbers := []float64{3, 5, 2, 9, 4}
	median := Median(numbers)
	fmt.Printf("Median: %.2f\n", median)
}

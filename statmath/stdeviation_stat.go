package statmath

import (
	"math"
)

func StandardDeviation(numbers []int) float64 {
	avg := AverageNum(numbers)
	variance := VarianceInt(numbers, avg)
	return math.Sqrt(float64(variance))
}

func VarianceInt(numbers []int, mean int) int {
	sum := 0
	for _, num := range numbers {
		sum += int(math.Pow(float64(num-mean), 2))
	}
	return sum / len(numbers)
}

func AverageNum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum / len(numbers)
}

package statmath

func Average(data []int) float64 {
	var sum int
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

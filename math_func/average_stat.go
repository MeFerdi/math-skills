package main

import "fmt"

func Average(data []int) float64 {
	var sum int
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

func main() {
	set_val := []int{12, 13, 74, 123, 11234}

	fmt.Println(Average(set_val))
}

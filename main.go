package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	statmath "statmath/statmath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}

	filePath := os.Args[1]
	data, err := readDataFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	average := statmath.Average(data)
	median := statmath.Median(data)
	variance := statmath.Variance(data)
	standarddeviation := statmath.StandardDeviation(data)
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard deviation: %.0f\n", standarddeviation)
}

func readDataFromFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

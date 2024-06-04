package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	statmath "statmath/statmath"
)

const maxValueThreshold = 100000

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}

	// Get the file path from the command-line argument
	filePath := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not found.")
		return
	}

	// Read data from the file
	data, err := readDataFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	// Check if the data is empty
	if len(data) == 0 {
		fmt.Println("No data found in the file.")
		return
	}

	// Check if any value exceeds the maximum threshold
	for _, value := range data {
		if value > maxValueThreshold {
			fmt.Println("Error: Value exceeds the maximum threshold.")
			return
		}
	}

	// Calculate and print the statistics
	average := math.Round(statmath.Average(data))
	median := math.Round(statmath.Median(data))
	variance := math.Round(statmath.Variance(data))
	standarddeviation := math.Round(statmath.StandardDeviation(data))
	fmt.Printf("Average: %d\n", int(average))
	fmt.Printf("Median: %d\n", int(median))
	fmt.Printf("Variance: %d\n", int(variance))
	fmt.Printf("Standard deviation: %d\n", int(standarddeviation))
}

func readDataFromFile(filePath string) ([]float64, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize an empty slice to store the data
	var data []float64

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the line from the scanner
		line := scanner.Text()

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Parse the value from the line
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}

		// Add the value to the data slice
		data = append(data, value)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Return the data
	return data, nil
}

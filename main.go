package main

import (
	"bufio"
	"fmt"
	statmath "math_func/statmath"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
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
	fmt.Printf("Average: %.0f\n", average)
}

func readDataFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
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

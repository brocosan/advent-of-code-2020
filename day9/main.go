package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var preambleLength = 25

func getInputData() []int {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	inputData := []int{}
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		inputData = append(inputData, number)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return inputData
}

func getInvalidNumber(inputData []int) int {
	numberRange := []int{}
mainLoop:
	for _, number := range inputData {
		// Build the range with the first N numbers
		if len(numberRange) < preambleLength {
			numberRange = append(numberRange, number)
			continue
		}

		// Search if this number is the sum of two numbers in the range
		for i, n := range numberRange {
			for j := i + 1; j < preambleLength; j++ {
				// We found a sum, this number is valid
				if n+numberRange[j] == number {
					// Add the number to the range and remove the first one
					numberRange = append(numberRange, number)
					numberRange = numberRange[1:]
					continue mainLoop
				}
			}
		}

		// We didn't find a sum, this number is invalid
		return number
	}
	return -1
}

func getEncryptionWeekness(inputData []int, targetNumber int) int {
	sum := 0
	rangeSum := []int{}
	for _, number := range inputData {
		// Ignore the invalid number
		if number == targetNumber {
			sum = 0
			rangeSum = []int{}
			continue
		}

		// Add sum with previous numbers
		sum += number
		rangeSum = append(rangeSum, number)

		// The sum is greater than the target, we remove the first number from the range
		if sum > targetNumber {
			for sum > targetNumber {
				sum -= rangeSum[0]
				rangeSum = rangeSum[1:]
			}
		}

		// The sum is the target number
		if sum == targetNumber {
			// Sort the slice and add the smallest and highest element
			sort.Ints(rangeSum)
			return rangeSum[0] + rangeSum[len(rangeSum)-1]
		}
	}

	return -1
}

func main() {
	inputData := getInputData()
	invalidNumber := getInvalidNumber(inputData)
	fmt.Printf("first invalid number: %d\n", invalidNumber)
	fmt.Printf("encryption weekness: %d\n", getEncryptionWeekness(inputData, invalidNumber))
}

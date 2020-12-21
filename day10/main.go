package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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
	sort.Ints(inputData)
	return inputData
}

func part1(inputData []int) {
	difference1, difference3 := 0, 1
	previousAdapter := 0
	for _, adapter := range inputData {
		switch adapter - previousAdapter {
		case 1:
			difference1++
		case 3:
			difference3++
		}
		previousAdapter = adapter
	}
	fmt.Printf("differences of 1: %d, differences of 3: %d \n", difference1, difference3)
	fmt.Printf("puzzle multiplication: %d \n", difference1*difference3)
}

func part2(inputData []int, startIndex int) int {
	inputData = append([]int{0}, inputData...)
	length := len(inputData)
	solutions := make([]int, length)
	solutions[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		solutions[i] = solutions[i+1]
		if i+2 < length && inputData[i+2]-inputData[i] <= 3 {
			solutions[i] += solutions[i+2]
		}
		if i+3 < length && inputData[i+3]-inputData[i] <= 3 {
			solutions[i] += solutions[i+3]
		}
	}
	return solutions[0]
}

func main() {
	inputData := getInputData()
	part1(inputData)
	fmt.Printf("number of solutions: %d \n", part2(inputData, 0))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() [][]string {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	inputData := [][]string{}
	for scanner.Scan() {
		seats := []string{}
		for _, seat := range scanner.Text() {
			seats = append(seats, string(seat))
		}
		inputData = append(inputData, seats)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return inputData
}

func applyRules(inputData [][]string) bool {
	changed := false
	seats := make([][]string, len(inputData))
	for i := range inputData {
		seats[i] = make([]string, len(inputData[i]))
		copy(seats[i], inputData[i])
	}
	for i, line := range inputData {
		for j, seat := range line {
			nbOccupiedSeats := countAdjacentOccupiedSeats(seats, i, j)
			switch seat {
			case "L":
				if nbOccupiedSeats == 0 {
					inputData[i][j] = "#"
					changed = true
				}
			case "#":
				if nbOccupiedSeats >= 4 {
					inputData[i][j] = "L"
					changed = true
				}
			}
		}
	}
	return changed
}

func getAdjacentSeats(seats [][]string, row, column int) []string {
	adjacent := []string{}
	height := len(seats)
	width := len(seats[0])

	up := row-1 >= 0
	down := row+1 < height
	left := column-1 >= 0
	right := column+1 < width

	// Up
	if up {
		adjacent = append(adjacent, seats[row-1][column])
	}
	// Down
	if down {
		adjacent = append(adjacent, seats[row+1][column])
	}
	// Left
	if left {
		adjacent = append(adjacent, seats[row][column-1])
	}
	// Right
	if right {
		adjacent = append(adjacent, seats[row][column+1])
	}
	// Up left
	if up && left {
		adjacent = append(adjacent, seats[row-1][column-1])
	}
	// Up right
	if up && right {
		adjacent = append(adjacent, seats[row-1][column+1])
	}
	// Down left
	if down && left {
		adjacent = append(adjacent, seats[row+1][column-1])
	}
	// Down right
	if down && right {
		adjacent = append(adjacent, seats[row+1][column+1])
	}

	return adjacent
}

func countAdjacentOccupiedSeats(seats [][]string, row, column int) int {
	count := 0
	for _, v := range getAdjacentSeats(seats, row, column) {
		if v == "#" {
			count++
		}
	}
	return count
}

func countOccupiedSeats(seats [][]string) int {
	count := 0
	for _, line := range seats {
		for _, seat := range line {
			if seat == "#" {
				count++
			}
		}
	}
	return count
}

func printSeats(inputData [][]string) {
	for _, line := range inputData {
		fmt.Println(line)
	}
	fmt.Println()
}

func main() {
	inputData := getInputData()
	for applyRules(inputData) {
		// printSeats(inputData)
	}
	fmt.Printf("number of occupied seats: %d\n", countOccupiedSeats(inputData))
}

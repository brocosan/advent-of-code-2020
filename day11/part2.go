package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputData = [][]string{}
var height = 0
var width = 0
var directions = []struct{ x, y int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func generateGrid() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
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
	height = len(inputData)
	width = len(inputData[0])
}

func applyRules() bool {
	changed := false
	seats := make([][]string, height)
	for i := range inputData {
		seats[i] = make([]string, width)
		copy(seats[i], inputData[i])
	}
	for i, line := range inputData {
		for j, seat := range line {
			switch seat {
			case "L":
				if canTakeSeat(seats, i, j) {
					inputData[i][j] = "#"
					changed = true
				}
			case "#":
				if leaveSeat(seats, i, j) {
					inputData[i][j] = "L"
					changed = true
				}
			}
		}
	}
	return changed
}

func canTakeSeat(seats [][]string, row, column int) bool {
	for _, d := range directions {
		for x, y := row+d.x, column+d.y; isInRange(x, y); x, y = x+d.x, y+d.y {
			if seats[x][y] == "L" {
				break
			}
			if seats[x][y] == "#" {
				return false
			}
		}
	}
	return true
}

func leaveSeat(seats [][]string, row, column int) bool {
	var occupied int
	for _, d := range directions {
		for x, y := row+d.x, column+d.y; isInRange(x, y); x, y = x+d.x, y+d.y {
			if seats[x][y] == "L" {
				break
			}
			if seats[x][y] == "#" {
				occupied++
				break
			}
		}
	}
	return occupied >= 5
}

func isInRange(row, column int) bool {
	return row >= 0 && row < height && column >= 0 && column < width
}

func countOccupiedSeats() int {
	count := 0
	for _, line := range inputData {
		for _, seat := range line {
			if seat == "#" {
				count++
			}
		}
	}
	return count
}

func printSeats() {
	for _, line := range inputData {
		fmt.Println(line)
	}
	fmt.Println()
}

func main() {
	generateGrid()
	// applyRules()
	// applyRules()
	// applyRules()
	// printSeats()
	for applyRules() {
		// printSeats()
	}
	fmt.Printf("number of occupied seats: %d\n", countOccupiedSeats())
	// fmt.Println(inputData[2][1])
}

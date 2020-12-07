package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	fmt.Printf("Part 1 number of trees: %d\n", getTreesEncounters(3, 1))
}

func part2() {
	slope11 := getTreesEncounters(1, 1)
	slope31 := getTreesEncounters(3, 1)
	slope51 := getTreesEncounters(5, 1)
	slope71 := getTreesEncounters(7, 1)
	slope12 := getTreesEncounters(1, 2)
	fmt.Printf("Part 2 number of trees 1-1: %d\n", slope11)
	fmt.Printf("Part 2 number of trees 3-1: %d\n", slope31)
	fmt.Printf("Part 2 number of trees 5-1: %d\n", slope51)
	fmt.Printf("Part 2 number of trees 7-1: %d\n", slope71)
	fmt.Printf("Part 2 number of trees 1-2: %d\n", slope12)
	fmt.Printf("Part 2 multiply: %d\n", slope11*slope31*slope51*slope71*slope12)
}

func getTreesEncounters(movesRight, movesDown int) int {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	newPosition, treesCount, descendCount := 0, 0, 0
	hasMovedRight := false
	for scanner.Scan() {
		line := scanner.Text()
		if hasMovedRight {
			descendCount++
			if descendCount == movesDown {
				if string(line[newPosition]) == "#" {
					treesCount++
				}
				hasMovedRight = false
				descendCount = 0
			} else {
				continue
			}
		}
		length := len(line)
		newPosition += movesRight
		if newPosition >= length {
			newPosition -= length
		}
		hasMovedRight = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return treesCount
}

func main() {
	part1()
	part2()
}

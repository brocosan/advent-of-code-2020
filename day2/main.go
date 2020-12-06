package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	validPasswordCount := 0
	for scanner.Scan() {
		// firstSplit[1] == password
		firstSplit := strings.Split(scanner.Text(), ": ")

		// secondSplit[1] == letter
		secondSplit := strings.Split(firstSplit[0], " ")

		// policy[0] == min / policy[1] == max
		policy := strings.Split(secondSplit[0], "-")
		minCount, _ := strconv.Atoi(policy[0])
		maxCount, _ := strconv.Atoi(policy[1])

		// Count the letter occurrences in the password
		occurrences := strings.Count(firstSplit[1], secondSplit[1])

		if occurrences >= minCount && occurrences <= maxCount {
			validPasswordCount++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1 number of valid password: %d\n", validPasswordCount)
}

func part2() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	validPasswordCount := 0
	for scanner.Scan() {
		// firstSplit[1] == password
		firstSplit := strings.Split(scanner.Text(), ": ")

		// secondSplit[1] == letter
		secondSplit := strings.Split(firstSplit[0], " ")
		positions := strings.Split(secondSplit[0], "-")
		position1, _ := strconv.Atoi(positions[0])
		position2, _ := strconv.Atoi(positions[1])
		letterAtPosition1 := string(firstSplit[1][position1-1])
		letterAtPosition2 := string(firstSplit[1][position2-1])

		if letterAtPosition1 == letterAtPosition2 {
			continue
		}
		if letterAtPosition1 == secondSplit[1] || letterAtPosition2 == secondSplit[1] {
			validPasswordCount++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 2 number of valid password: %d\n", validPasswordCount)
}

func main() {
	part1()
	part2()
}

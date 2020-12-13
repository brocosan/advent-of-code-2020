package main

import (
	"bufio"
	"fmt"
	"os"
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
	sum := 0
	yesQuestions := make(map[rune]bool)
	for scanner.Scan() {
		line := scanner.Text()

		// New line: this is a new group
		if line == "" {
			sum += len(yesQuestions)
			yesQuestions = make(map[rune]bool)
			continue
		}

		for _, char := range line {
			yesQuestions[char] = true
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Check if we have a last group to check
	if yesQuestions != nil {
		sum += len(yesQuestions)
	}

	fmt.Printf("Part 1 - sum of each group: %d\n", sum)
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
	sum := 0
	numberOfPerson := 0
	yesQuestions := make(map[rune]int)
	for scanner.Scan() {
		line := scanner.Text()

		// New line: this is a new group
		if line == "" {
			for _, charCount := range yesQuestions {
				if charCount == numberOfPerson {
					sum++
				}
			}

			yesQuestions = make(map[rune]int)
			numberOfPerson = 0
			continue
		}

		numberOfPerson++
		for _, char := range line {
			yesQuestions[char]++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Check if we have a last group to check
	if yesQuestions != nil {
		for _, charCount := range yesQuestions {
			if charCount == numberOfPerson {
				sum++
			}
		}
	}

	fmt.Printf("Part 2 - sum of each group: %d\n", sum)
}

func main() {
	part1()
	part2()
}

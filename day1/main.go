package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getContentFromFile() []int {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the numbers
	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort the array
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	return numbers
}

func getSumTwoNumbers(numbers []int, goal int) (int, int) {
	for i, first := range numbers {
		if first > goal {
			continue
		}
		for _, second := range numbers[i+1:] {
			if first+second == goal {
				return first, second
			}
		}
	}
	return 0, 0
}

func getSumThreeNumbers(numbers []int, goal int) (int, int, int) {
	for i, first := range numbers {
		if first > goal {
			continue
		}
		for j, second := range numbers[i+1:] {
			if first+second > goal {
				continue
			}
			for _, third := range numbers[j+1:] {
				if first+second+third == goal {
					return first, second, third
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	numbers := getContentFromFile()

	// Part 1
	number1, number2 := getSumTwoNumbers(numbers, 2020)
	if number1 == 0 && number2 == 0 {
		panic("Numbers not found")
	}
	fmt.Printf("first: %d - second: %d / multiply: %d\n", number1, number2, number1*number2)

	// Part 2
	number1, number2, number3 := getSumThreeNumbers(numbers, 2020)
	if number1 == 0 && number2 == 0 && number3 == 0 {
		panic("Numbers not found")
	}
	fmt.Printf("first: %d - second: %d - third: %d / multiply: %d\n", number1, number2, number3, number1*number2*number3)
}

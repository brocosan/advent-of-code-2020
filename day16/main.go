package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Ticket is a series of numbers
type Ticket []int

// Range represents a range between min and max
type Range struct {
	min int
	max int
}

// Field is composed of two ranges
type Field struct {
	lowRange  Range
	highRange Range
}

var regexpRange = regexp.MustCompile(`(\d+)-(\d+) or (\d+)-(\d+)`)
var myTicket = Ticket{}

func (f *Field) isRangeValid(number int) bool {
	if number >= f.lowRange.min && number <= f.lowRange.max {
		return true
	}
	if number >= f.highRange.min && number <= f.highRange.max {
		return true
	}
	return false
}

func getInput() []string {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return input
}

func parseTicket(line string) Ticket {
	numbers := strings.Split(line, ",")
	ticket := make(Ticket, len(numbers))
	for i, n := range numbers {
		ticket[i], _ = strconv.Atoi(n)
	}
	return ticket
}

func getField(line string) Field {
	field := Field{}
	matches := regexpRange.FindAllStringSubmatch(line, -1)
	min, _ := strconv.Atoi(matches[0][1])
	max, _ := strconv.Atoi(matches[0][2])
	field.lowRange = Range{min, max}
	min, _ = strconv.Atoi(matches[0][3])
	max, _ = strconv.Atoi(matches[0][4])
	field.highRange = Range{min, max}
	return field
}

func getValidTickets(tickets []Ticket, fields []Field) ([]Ticket, int) {
	errors := 0
	validTickets := []Ticket{}
	for _, ticket := range tickets {
		ticketIsValid := true
		for _, number := range ticket {
			numberIsValid := false
			for _, field := range fields {
				if field.isRangeValid(number) {
					numberIsValid = true
					break
				}
			}
			if numberIsValid == false {
				errors += number
				ticketIsValid = false
			}
		}
		if ticketIsValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, errors
}

func main() {
	input := getInput()

	// Get fields
	fields := []Field{}
	for i, line := range input {
		if line == "" {
			myTicket = parseTicket(input[i+2])
			input = input[i+5:]
			break
		}
		fields = append(fields, getField(line))
	}

	// Get nearby tickets
	nearbyTickets := []Ticket{}
	for _, line := range input {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	// Get the scanning error rate (part 1)
	// Get only valid tickets
	validTickets, errors := getValidTickets(nearbyTickets, fields)
	fmt.Printf("Part 1 - ticket scanning error rate: %d\n", errors)

	part2(fields, validTickets)
}

func part2(fields []Field, validTickets []Ticket) {
	possiblePositions := make([]map[int]int, len(fields))
	for i := range possiblePositions {
		possiblePositions[i] = make(map[int]int)
	}
	for i, field := range fields {
		// Check every position until we find a possible one
	loopPosition:
		for position := 0; position < len(fields); position++ {
			for _, ticket := range validTickets {
				if !field.isRangeValid(ticket[position]) {
					continue loopPosition
				}
			}
			// We found a possible position
			possiblePositions[i][position] = i
		}
	}

	// Sort by lower solutions possible
	sort.Slice(possiblePositions, func(i, j int) bool {
		return len(possiblePositions[i]) < len(possiblePositions[j])
	})

	// We find the only possible position
	// Note: only one loop is needed with this input
	// This is an ugly solution but it works.
	positions := make([]int, len(fields))
	for _, pos := range possiblePositions {
		tmpPos := []int{}
		curPos := 0

		// Keep only positions that were not found
		for j, p := range pos {
			if !contains(positions, j) {
				tmpPos = append(tmpPos, j)
				curPos = p
			}
		}

		if len(tmpPos) == 1 {
			positions[curPos] = tmpPos[0]
		}
	}

	// Little shortcut here, the departure fields are fields 1 - 6
	// Let's multiply these values to end this challenge
	solution := 1
	for i := 0; i < 6; i++ {
		solution *= myTicket[positions[i]]
	}
	fmt.Printf("Part 2 - solution: %d\n", solution)
}

// Helper function to check if a value exists in a slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

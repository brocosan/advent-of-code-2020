package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var directions = map[int]string{0: "E", 1: "S", 2: "W", 3: "N"}

// Ship has a face (current direction) and a record of positions
type Ship struct {
	face      int
	positions map[string]int
}

func (s *Ship) doAction(action string, parameter int) {
	switch action {
	case "N", "S", "E", "W":
		s.positions[action] += parameter
	case "F":
		s.positions[directions[s.face]] += parameter
	case "L":
		s.turn(parameter * -1)
	case "R":
		s.turn(parameter)
	}
}

func (s *Ship) turn(angle int) {
	index := (angle/90 + s.face) % 4
	if index < 0 {
		index += 4
	}
	s.face = index
}

func (s *Ship) getManhattanDistance() int {
	northSouth := s.positions["N"] - s.positions["S"]
	if s.positions["N"] < s.positions["S"] {
		northSouth = s.positions["S"] - s.positions["N"]
	}
	eastWest := s.positions["E"] - s.positions["W"]
	if s.positions["E"] < s.positions["W"] {
		eastWest = s.positions["W"] - s.positions["E"]
	}
	return northSouth + eastWest
}

func main() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	ship := Ship{face: 0}
	ship.positions = map[string]int{"N": 0, "E": 0, "S": 0, "W": 0}
	matcher := regexp.MustCompile(`(N|S|E|W|L|R|F)(\d+)`)
	for scanner.Scan() {
		matches := matcher.FindAllStringSubmatch(scanner.Text(), -1)
		parameter, _ := strconv.Atoi(matches[0][2])
		ship.doAction(matches[0][1], parameter)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("distance: %d\n", ship.getManhattanDistance())
}

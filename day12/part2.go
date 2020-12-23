package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var directions = map[int]string{0: "E", 1: "S", 2: "W", 3: "N"}

// Ship has a waypoint and a record of positions
type Ship struct {
	waypoint  map[int]int
	positions map[string]int
}

func (s *Ship) doAction(action string, parameter int) {
	switch action {
	case "E":
		s.waypoint[0] += parameter
	case "S":
		s.waypoint[1] += parameter
	case "W":
		s.waypoint[2] += parameter
	case "N":
		s.waypoint[3] += parameter
	case "F":
		s.positions["E"] += s.waypoint[0] * parameter
		s.positions["S"] += s.waypoint[1] * parameter
		s.positions["W"] += s.waypoint[2] * parameter
		s.positions["N"] += s.waypoint[3] * parameter
	case "L":
		s.turn(parameter * -1)
	case "R":
		s.turn(parameter)
	}
}

func (s *Ship) turn(angle int) {
	newWaypoint := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}
	for i, point := range s.waypoint {
		index := (angle/90 + i) % 4
		if index < 0 {
			index += 4
		}
		newWaypoint[index] = point
	}
	s.waypoint = newWaypoint
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
	ship := Ship{}
	// The waypoint starts 10 units east and 1 unit north
	ship.waypoint = map[int]int{0: 10, 1: 0, 2: 0, 3: 1}
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

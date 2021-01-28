package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getActiveCubes() []string {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)

	activeCubes := []string{}
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		for y, v := range line {
			if string(v) == "#" {
				activeCubes = append(activeCubes, fmt.Sprintf("%d,%d,%d,%d", x, y, 0, 0))
			}
		}
		x++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return activeCubes
}

func getNeighbors(activeCubes []string, c string) []string {
	cube := getFields(c)
	neighbors := []string{}
	for ww := -1; ww <= 1; ww++ {
		for zz := -1; zz <= 1; zz++ {
			for xx := -1; xx <= 1; xx++ {
				for yy := -1; yy <= 1; yy++ {
					activeCube := fmt.Sprintf("%d,%d,%d,%d", xx+cube[0], yy+cube[1], zz+cube[2], ww+cube[3])
					// Skip this cube
					if activeCube == c {
						continue
					}
					neighbors = append(neighbors, activeCube)
				}
			}
		}
	}
	return neighbors
}

func doCycle(activeCubes []string) []string {
	state := []string{}
	inactiveCubesToCheck := map[string]int{}
	// Get new state (new active cubes)
	for _, activeCube := range activeCubes {
		// Get all neighbors for a cube
		neighbors := getNeighbors(activeCubes, activeCube)

		activeNeighborCount := 0
		for _, neighbor := range neighbors {
			if contains(activeCubes, neighbor) {
				activeNeighborCount++
			} else {
				inactiveCubesToCheck[neighbor]++
			}
		}

		// This cube stays active
		if activeNeighborCount == 2 || activeNeighborCount == 3 {
			state = append(state, activeCube)
		}
	}

	// Let's check the inactive cube that could change status
	for cube, activeNeighbors := range inactiveCubesToCheck {
		// This cube becomes active
		if activeNeighbors == 3 {
			state = append(state, cube)
		}
	}

	return state
}

func getFields(str string) []int {
	s := strings.Split(str, ",")
	fields := make([]int, len(s))
	for i, v := range s {
		fields[i], _ = strconv.Atoi(v)
	}
	return fields
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	// Parse the input to get the active cubes
	activeCubes := getActiveCubes()

	// Solve the problem
	for i := 0; i < 6; i++ {
		activeCubes = doCycle(activeCubes)
	}
	fmt.Printf("Number of active cubes: %d\n", len(activeCubes))
}

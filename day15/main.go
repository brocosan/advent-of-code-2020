package main

import "fmt"

var input = []int{14, 8, 16, 0, 1, 17}

func findNumber(limit int) int {
	lastSeen := map[int]int{}
	for i, v := range input {
		lastSeen[v] = i + 1
	}
	lastNumber := 0
	for i := len(input) + 2; i < limit+1; i++ {
		lastRound, seen := lastSeen[lastNumber]
		numberSpoken := 0
		if seen {
			numberSpoken = i - lastRound - 1
		}
		lastSeen[lastNumber] = i - 1
		lastNumber = numberSpoken
	}
	return lastNumber
}

func main() {
	fmt.Printf("the 2020th number spoken is %d\n", findNumber(2020))
	fmt.Printf("the 30000000th number spoken is %d\n", findNumber(30000000))
}

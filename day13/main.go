package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	timestamp, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := []int{}
	for _, ID := range strings.Split(scanner.Text(), ",") {
		if ID == "x" {
			continue
		}
		busID, _ := strconv.Atoi(ID)
		buses = append(buses, busID)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	wait := -1
	for {
		wait++
		for _, bus := range buses {
			if timestamp%bus == 0 {
				fmt.Printf("waited for: %d / bus ID: %d / multiply: %d\n", wait, bus, wait*bus)
				return
			}
		}
		timestamp++
	}
}

func part2() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	scanner.Scan()
	buses := map[int]int{}
	for i, ID := range strings.Split(scanner.Text(), ",") {
		if ID == "x" {
			continue
		}
		busID, _ := strconv.Atoi(ID)
		buses[i] = busID
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	timestamp := 0
	step := 1
	for i, bus := range buses {
		for (i+timestamp)%bus != 0 {
			timestamp += step
		}
		step *= bus
	}
	fmt.Printf("earliest timestamp: %d\n", timestamp)
}

func main() {
	part1()
	part2()
}

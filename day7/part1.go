package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var colorsReverse = map[string][]string{}
var visited = map[string]bool{}

func main() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " bags contain ")
		for _, bags := range strings.Split(strings.TrimSuffix(split[1], "."), ", ") {
			if bags == "no other bags" {
				continue
			}
			bags = strings.ReplaceAll(bags, " bags", "")
			bags = strings.ReplaceAll(bags, " bag", "")
			bag := strings.SplitN(bags, " ", 2)
			colorsReverse[bag[1]] = append(colorsReverse[bag[1]], split[0])
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	search("shiny gold", 0)
	fmt.Printf("bag colors that can contain at least one shiny gold bag: %d\n", len(visited))
}

func search(target string, x int) {
	colors, ok := colorsReverse[target]
	if ok {
		for _, color := range colors {
			visited[color] = true
			search(color, x+1)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Rule represents how many bags can be contained
type Rule struct {
	number int
	color  string
}

func countBags(rules map[string][]Rule, target string) int {
	count := 0
	for _, rule := range rules[target] {
		count += rule.number * (1 + countBags(rules, rule.color))
	}
	return count
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
	rules := map[string][]Rule{}
	matcher := regexp.MustCompile(`(\d+)\s(\w+\s\w+)\sbags?`)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " bags contain ")
		matches := matcher.FindAllStringSubmatch(split[1], -1)
		if matches == nil {
			continue
		}
		rules[split[0]] = make([]Rule, len(matches))
		for i, bag := range matches {
			rule := Rule{color: bag[2]}
			rule.number, _ = strconv.Atoi(bag[1])
			rules[split[0]][i] = rule
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Number of bags required: %d\n", countBags(rules, "shiny gold"))
}

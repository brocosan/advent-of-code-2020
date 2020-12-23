package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Code represents the series of instructions
type Code struct {
	instructions []Instruction
	accumulator  int
}

// Instruction represents of line of code
type Instruction struct {
	operation string
	argument  int
}

func (c *Code) execute() bool {
	executed := make(map[int]bool, len(c.instructions))
	for i := 0; i < len(c.instructions); i++ {
		instruction := c.instructions[i]
		if executed[i] {
			// It is an infinite loop
			return false
		}
		executed[i] = true
		c.instructions[i] = instruction
		if instruction.operation == "nop" {
			continue
		}
		if instruction.operation == "acc" {
			c.accumulator += instruction.argument
		}
		if instruction.operation == "jmp" {
			i += instruction.argument - 1
		}
	}
	return true
}

func (c *Code) debug() *Code {
	for i := 0; i < len(c.instructions); i++ {
		instruction := c.instructions[i]

		if instruction.argument == 0 {
			continue
		}

		if instruction.operation == "nop" {
			instruction.operation = "jmp"
			c.instructions[i] = instruction
			c.accumulator = 0
			if c.execute() {
				return c
			}
			instruction.operation = "nop"
			c.instructions[i] = instruction
			continue
		}
		if instruction.operation == "jmp" {
			instruction.operation = "nop"
			c.instructions[i] = instruction
			c.accumulator = 0
			if c.execute() {
				return c
			}
			instruction.operation = "jmp"
			c.instructions[i] = instruction
			continue
		}
	}
	return nil
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
	matcher := regexp.MustCompile(`(nop|acc|jmp) (\+|-)(\d+)`)
	code := Code{}
	for scanner.Scan() {
		matches := matcher.FindAllStringSubmatch(scanner.Text(), -1)
		if matches == nil {
			continue
		}
		instruction := Instruction{}
		instruction.operation = matches[0][1]
		argument, _ := strconv.Atoi(matches[0][3])
		if matches[0][2] == "-" {
			argument *= -1
		}
		instruction.argument = argument
		code.instructions = append(code.instructions, instruction)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	code.execute()
	fmt.Printf("accumulator value with infinite loop: %d\n", code.accumulator)

	debugged := code.debug()
	fmt.Printf("accumulator value after debug: %d\n", debugged.accumulator)
}

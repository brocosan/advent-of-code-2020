package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memory map[int]uint64
var memoryV2 map[int64]uint64
var mask string

func assignMemory(index, value int) {
	newValue, _ := strconv.ParseUint(applyMask(getBinary(value)), 2, 64)
	memory[index] = newValue
}

func applyMask(value string) string {
	newValue := []rune(value)
	for i, bit := range mask {
		if bit == 'X' {
			continue
		}
		newValue[i] = bit
	}
	return string(newValue)
}

func assignMemoryV2(index, value int) {
	indexesOfX := []int{}
	for i, bit := range mask {
		if bit == 'X' {
			indexesOfX = append(indexesOfX, i)
		}
	}
	indexMask := applyMaskV2(index)
	numberOfX := strings.Count(mask, "X")
	format := "%0" + strconv.Itoa(numberOfX) + "s"
	for i := float64(0); i < math.Pow(float64(2), float64(numberOfX)); i++ {
		replace := []rune(fmt.Sprintf(format, strconv.FormatInt(int64(i), 2)))
		newIndex := []rune(indexMask)
		for j, ix := range indexesOfX {
			newIndex[ix] = replace[j]
		}
		indexInt, _ := strconv.ParseInt(string(newIndex), 2, 64)
		memoryV2[indexInt] = uint64(value)
	}
}

func applyMaskV2(value int) string {
	newValue := []rune(getBinary(value))
	for i, bit := range mask {
		if bit == '0' {
			continue
		}
		newValue[i] = bit
	}
	return string(newValue)
}

func getBinary(value int) string {
	binaryValue := strconv.FormatInt(int64(value), 2)
	return fmt.Sprintf("%036s", binaryValue)
}

func main() {
	memory = make(map[int]uint64)
	memoryV2 = make(map[int64]uint64)
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	matcherMask := regexp.MustCompile(`mask = ([01X]+)`)
	matcherMem := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			matches := matcherMask.FindAllStringSubmatch(line, -1)
			mask = matches[0][1]
		} else if strings.HasPrefix(line, "mem") {
			matches := matcherMem.FindAllStringSubmatch(line, -1)
			index, _ := strconv.Atoi(matches[0][1])
			value, _ := strconv.Atoi(matches[0][2])
			assignMemory(index, value)
			assignMemoryV2(index, value)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part 1
	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("part1 sum: %d\n", sum)

	// Part 2
	sum = 0
	for _, v := range memoryV2 {
		sum += v
	}
	fmt.Printf("part2 sum: %d\n", sum)
}

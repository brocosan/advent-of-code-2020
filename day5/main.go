package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRow(chars string) int {
	start := 0
	end := 127
	for _, c := range chars {
		if c == 'F' {
			end = ((end - start) / 2) + start
		} else {
			start = ((end - start) / 2) + start + 1
		}
	}
	return start
}

func getColumn(chars string) int {
	start := 0
	end := 7
	for _, c := range chars {
		if c == 'L' {
			end = ((end - start) / 2) + start
		} else {
			start = ((end - start) / 2) + start + 1
		}
	}
	return start
}

func getID(row, column int) int {
	return row*8 + column
}

func main() {
	// Test examples
	passes := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	for _, v := range passes {
		row := getRow(v[0:7])
		column := getColumn(v[7:10])
		fmt.Printf("%v: row %d, column %d, seat ID %d\n", v, row, column, getID(row, column))
	}

	highest := 0

	// Open the file
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file line by line to get the data
	scanner := bufio.NewScanner(f)
	takenSeats := [128][8]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		row := getRow(line[0:7])
		column := getColumn(line[7:10])
		ID := getID(row, column)
		if ID > highest {
			highest = ID
		}
		takenSeats[row][column] = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("highest seatID: %d\n", highest)

	for indexRow, row := range takenSeats {
		for indexColumn, seat := range row {
			if !seat && indexColumn < 7 {
				next := takenSeats[indexRow][indexColumn+1]
				if next {
					fmt.Printf("my seat: row %d column %d ID %d\n", indexRow, indexColumn, getID(indexRow, indexColumn))
				}
			}
		}
	}
}

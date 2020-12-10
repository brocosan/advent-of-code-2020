package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport stores a passport's fields
type Passport map[string]string

var validPassports1 = 0
var validPassports2 = 0

func checkPassport(passport Passport) {
	if isValidPassport1(passport) {
		validPassports1++
	}
	if isValidPassport2(passport) {
		validPassports2++
	}
}

func isValidPassport1(passport Passport) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range requiredFields {
		_, ok := passport[field]
		if !ok {
			return false
		}
	}
	return true
}

func isValidPassport2(passport Passport) bool {
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	byr, err = strconv.Atoi(passport["iyr"])
	if err != nil || byr < 2010 || byr > 2020 {
		return false
	}

	byr, err = strconv.Atoi(passport["eyr"])
	if err != nil || byr < 2020 || byr > 2030 {
		return false
	}

	hgt, ok := passport["hgt"]
	if !ok || !isValidHeight(hgt) {
		return false
	}

	hcl, ok := passport["hcl"]
	if !ok || !isValidColor(hcl) {
		return false
	}
	ecl, ok := passport["ecl"]
	if !ok || !isValidEyeColor(ecl) {
		return false
	}
	pid, ok := passport["pid"]
	if !ok || !isValidPassportID(pid) {
		return false
	}

	return true
}

func isValidHeight(height string) bool {
	if strings.HasSuffix(height, "cm") {
		height = strings.TrimSuffix(height, "cm")
		number, err := strconv.Atoi(height)
		return err == nil && number >= 150 && number <= 193
	}
	if strings.HasSuffix(height, "in") {
		height = strings.TrimSuffix(height, "in")
		number, err := strconv.Atoi(height)
		return err == nil && number >= 59 && number <= 76
	}
	return false
}

func isValidColor(color string) bool {
	return regexp.MustCompile(`^#[0-9-a-f]{6}$`).MatchString(color)
}

func isValidEyeColor(color string) bool {
	switch color {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func isValidPassportID(pid string) bool {
	return regexp.MustCompile(`^\d{9}$`).MatchString(pid)
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
	var passport = make(Passport)
	for scanner.Scan() {
		line := scanner.Text()

		// New line: this is a new passport
		if line == "" {
			checkPassport(passport)
			passport = make(Passport)
			continue
		}

		// Passport line
		fields := strings.Split(line, " ")
		for _, field := range fields {
			item := strings.Split(field, ":")
			passport[item[0]] = item[1]
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Check if we have a last passport to check
	if passport != nil {
		checkPassport(passport)
	}

	fmt.Printf("Part1 - number of valid passports: %d\n", validPassports1)
	fmt.Printf("Part2 - number of valid passports: %d\n", validPassports2)
}

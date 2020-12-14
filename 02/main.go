package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input string) int {
	s := strings.Split(input, "\n")
	total := 0
	for _, line := range s {
		split := strings.Split(line, ": ")
		meta, pwd := split[0], split[1]
		split = strings.Split(meta, " ")
		minmax, char := split[0], split[1]
		split = strings.Split(minmax, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		count := strings.Count(pwd, char)
		if count >= min && count <= max {
			total += 1
		}
	}
	return total
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	s := strings.Split(input, "\n")
	total := 0
	for _, line := range s {
		split := strings.Split(line, ": ")
		meta, pwd := split[0], split[1]
		split = strings.Split(meta, " ")
		minmax, char := split[0], split[1]
		split = strings.Split(minmax, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		if (string(pwd[min-1])) == char != (string(pwd[max-1]) == char) {
			total += 1
		}
	}
	return total
}

func main() {
	file := utils.Input(2020, 2)
	start := time.Now()
	fmt.Println("Part 1: ", Part1(file), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(file), "Time", time.Since(start))
}

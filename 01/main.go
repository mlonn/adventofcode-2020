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
	m := make(map[int]int)
	for _, expense := range s {
		i ,_ := strconv.Atoi(expense)
		complement := 2020 - i
		_, found := m[complement]
		if found {
			return i * complement
		}
		m[i] = i
	}
	panic("Not found")
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	s := strings.Split(input, "\n")
	m := make(map[int]int)
	for _, expense1 := range s {
		i1 ,_ := strconv.Atoi(expense1)
		for _, expense2 := range s {
			i2 ,_ := strconv.Atoi(expense2)
			complement := 2020 - i1 - i2
			_, found := m[complement]
			if found {
				return i1*i2*complement
			}
			m[i2] = i2
		}
		m[i1] = i1
	}
	panic("Not found")
}

func main() {
	input := utils.Input(2020,1)
	start := time.Now()
	fmt.Println("Part 1: " , Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " , Part2(input),"Time", time.Since(start))
}
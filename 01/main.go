package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


// Part1 Part 1 of puzzle
func Part1(input string) string{
	s := strings.Split(input, "\n")
	m := make(map[int]int)

	for _, expense := range s {
		i ,_ := strconv.Atoi(expense)
		complement := 2020 - i
		_, found := m[complement]
		if found {
			return "Answer: " + strconv.Itoa(i * complement)
		}
		m[i] = i
	}
	return "Not found"
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	s := strings.Split(input, "\n")
	m := make(map[int]int)
	for _, expense1 := range s {
		i1 ,_ := strconv.Atoi(expense1)
		for _, expense2 := range s {
			i2 ,_ := strconv.Atoi(expense2)
			complement := 2020 - i1 - i2
			_, found := m[complement]
			if found {
				return "Answer: " + strconv.Itoa(i1*i2*complement)
			}
			m[i2] = i2
		}
		m[i1] = i1
	}
	return "Not found"
}

func main() {
	bytes, _ := ioutil.ReadFile("01/input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
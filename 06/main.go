package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)


// Part1 Part 1 of puzzle
func Part1(input string) string {
	sum := 0
	groups := strings.Split(input,"\n\n")
	for _, group := range groups {
		answers := strings.Split(group,"\n")
		m := make(map[int32]bool)
		for _, answer := range answers {
			for _, question := range answer {
				m[question] = true
			}

		}
		sum+= len(m)
	}
	return "Answer " + strconv.Itoa(sum)
}

// Part2 Part2 of puzzle
func Part2(input string) string {

	sum := 0
	groups := strings.Split(input,"\n\n")
	for _, group := range groups {
		answers := strings.Split(group,"\n")
		m := make(map[int32]int)
		for _, answer := range answers {
			for _, question := range answer {
				m[question] += 1
			}

		}
		for _, yes := range m {
			if yes == len(answers) {
				sum += 1
			}
		}

	}
	return "Answer " + strconv.Itoa(sum)
}

func main() {
	start := time.Now()
	bytes, _ := ioutil.ReadFile("06/input.txt")
	input := string(bytes)
	fmt.Println("Read file in: ", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(input),"Time", time.Since(start))
}

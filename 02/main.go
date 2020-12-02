package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)


// Part1 Part 1 of puzzle
func Part1(input string) string {
	s := strings.Split(input, "\n")
	total:= 0
	for _, line := range s {
		re, _ := regexp.Compile(`(\d*)-(\d*) (.): (.*)`)

		matches := re.FindAllStringSubmatch(line,-1)
		match := matches[0]
		min,_ := strconv.Atoi(match[1])
		max,_ := strconv.Atoi(match[2])
		char := match[3]
		pwd := match[4]
		count := strings.Count(pwd, char)
		if count >= min && count <= max {
			total += 1
		}
	}
	return "Answer " + strconv.Itoa(total)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	s := strings.Split(input, "\n")
	total:= 0
	for _, line := range s {
		re, _ := regexp.Compile(`(\d*)-(\d*) (.): (.*)`)
		matches := re.FindAllStringSubmatch(line,-1)
		match := matches[0]
		min,_ := strconv.Atoi(match[1])
		max,_ := strconv.Atoi(match[2])
		char := match[3]
		pwd := match[4]
		if (string(pwd[min-1])) == char != (string(pwd[max-1]) == char) {
			total += 1
		}
	}
	return "Answer " + strconv.Itoa(total)
}

func main() {
	bytes, _ := ioutil.ReadFile("02/input.txt")
	start1 := time.Now()
	fmt.Println("Part 1: " + Part1(string(bytes)), "Time", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Part 2: " + Part2(string(bytes)),"Time", time.Since(start2))
}
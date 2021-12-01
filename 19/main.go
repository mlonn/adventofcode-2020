package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type rule struct {
	rules []rule
	r     string
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	rulePart := strings.Split(parts[0], "\n")
	//messages := parts[1]
	r := make(map[int]string)
	for _, rule := range rulePart {
		split := strings.Split(rule, ": ")
		n, _ := strconv.Atoi(split[0])
		r[n] = split[1]
	}
	//rules := make(map[int]rule)
	for _, rule := range r {
		split := strings.Split(rule, " | ")

		fmt.Println(split)
	}
	return -1
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	return -1
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 19)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: "+strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: "+strconv.Itoa(Part2(input)), "\t", time.Since(start))
}

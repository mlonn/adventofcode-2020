package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func parseInput(input string) (int, map[int]int, int) {
	lines := strings.Split(input, "\n")
	timestamp, _ := strconv.Atoi(lines[0])
	busses := make(map[int]int, 0)
	maxBus := 0
	for i, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}
		busNumber, _ := strconv.Atoi(bus)
		if busNumber > maxBus {
			maxBus = busNumber
		}
		busses[busNumber] = i
	}
	return timestamp, busses, maxBus
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	timestamp, busses, _ := parseInput(input)
	minWait := math.MaxInt32
	minBus := 0
	for bus := range busses {
		wait := 0
		for i := 0; i <= timestamp/bus; i++ {
			wait += bus
		}
		if wait < minWait {
			minWait = wait
			minBus = bus
		}
	}
	return (minWait - timestamp) * minBus
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	_, busses, maxBus := parseInput(input)
	found := false

	t := 0
	for !found {
		t += maxBus
		works := true
		for bus := range busses {
			t0 := t - busses[maxBus]
			if (t0+busses[bus])%bus != 0 {
				works = false
			}
		}
		found = works
	}
	return t
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 13)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: "+strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: "+strconv.Itoa(Part2(input)), "\t", time.Since(start))
}

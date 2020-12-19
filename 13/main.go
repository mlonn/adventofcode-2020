package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Bus struct {
	index, interval int
}

func parseInput(input string) (int, []Bus) {
	lines := strings.Split(input, "\n")
	timestamp, _ := strconv.Atoi(lines[0])
	busses := make([]Bus, 0)
	for i, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}
		busNumber, _ := strconv.Atoi(bus)
		busses = append(busses, Bus{index: i, interval: busNumber})
	}
	return timestamp, busses
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	timestamp, busses := parseInput(input)
	minWait := math.MaxInt8
	minBus := 0
	for _, bus := range busses {
		wait := 0
		for i := 0; i <= timestamp/bus.interval; i++ {
			wait += bus.interval
		}
		if wait < minWait {
			minWait = wait
			minBus = bus.interval
		}

	}
	return (minWait - timestamp) * minBus
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	_, busses := parseInput(input)
	var alignedBusses []Bus
	for _, bus := range busses {
		start := (bus.interval - bus.index) % bus.interval
		if start < 0 {
			start += bus.interval
		}
		alignedBusses = append(alignedBusses, Bus{
			index:    start,
			interval: bus.interval,
		})
	}
	step := 1
	t := 0
	for _, bus := range alignedBusses {
		for t%bus.interval != bus.index {
			t += step
		}
		step *= bus.interval
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

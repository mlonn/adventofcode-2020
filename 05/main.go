package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)



func findRow(inputs string, lower int,upper int)  int {
	if upper == lower {
		return upper
	}
	if len(inputs) == 0 {
		return upper
	}
	current, rest := inputs[0], inputs[1:]
	middle := (lower + upper) / 2
	if current == 'F' {
		return findRow(rest, lower, middle)
	}
	if current == 'B' {
		return findRow(rest, middle, upper)

	}
	return upper
}

func findSeat(inputs string, lower int,upper int)  int {
	if upper == lower {
		return upper
	}
	if len(inputs) == 0 {
		return upper
	}
	current, rest := inputs[0], inputs[1:]
	middle := (lower + upper) / 2
	if current == 'L' {
		return findSeat(rest, lower, middle)
	}
	if current == 'R' {
		return findSeat(rest, middle, upper)

	}
	return upper
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	max:= 0
	boardingPasses := strings.Split(input,"\n")
	for _, boardingPass := range boardingPasses {
		row, seat := boardingPass[0:7], boardingPass[7:]
		id := findRow(row, 0,127) * 8 + findSeat(seat, 0,7)
		if id > max {
			max = id
		}
	}
	return "Answer " + strconv.Itoa(max)
}

// Part2 Part2 of puzzle
func Part2(input string) string {


	boardingPasses := strings.Split(input,"\n")
	ids := make([]int, 0)
	for _, boardingPass := range boardingPasses {
		row, seat := boardingPass[0:7], boardingPass[7:]
		id := findRow(row, 0,127) * 8 + findSeat(seat, 0,7)
		ids = append(ids,id)
	}
	sort.Ints(ids)
	for i, id := range ids {
		if i < len(ids) -1 {
			if  ids[i+1] - id  > 1{
				return "Answer " + strconv.Itoa((ids[i+1]+id)/2)
			}
		}

	}
	return "Not found"
}

func main() {
	file := utils.Input(2020,5)
	start := time.Now()
	fmt.Println("Part 1: " + Part1(file), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(file),"Time", time.Since(start))
}

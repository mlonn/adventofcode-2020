package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)



// Part1 Part 1 of puzzle
func Part1(input string) int {
	size := 25
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if i < size {
			continue
		}
		
		preamble := lines[i-size :i]
		current, _ := strconv.Atoi(line)
		found := false

		for _, first := range preamble {
			for _, second := range preamble {
				x, _ := strconv.Atoi(first)
				y, _ := strconv.Atoi(second)
				if x + y == current {
					found = true
				}
			}
		}
		if !found {
			return current
		}
	}
	panic("Not found")
}

// Part2 Part2 of puzzle
func Part2(input string, weakness int) int {
	lines := strings.Split(input, "\n")
	for i, _ := range lines {
		sum :=0
		list := make([]int,0)
		for j := i;j < len(lines); j++ {
			current, _ := strconv.Atoi(lines[j])
			list = append(list,current)
			if sum == weakness {
				sort.Ints(list)
				return list[0]+list[len(list)-1]
			}
			sum += current
		}
	}
	panic("Not found")
}

func main() {
	start := time.Now()
	input := utils.Input(2020,9)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	weakness := Part1(input)
	fmt.Println("Part 1: " + strconv.Itoa(weakness), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + strconv.Itoa(Part2(input, weakness)),"\t", time.Since(start))
}

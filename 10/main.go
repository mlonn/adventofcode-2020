package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)



// Part1 Part 1 of puzzle
func Part1(input string) int {
	three := 0
	one := 0
	lines := strings.Split(input, "\n")
	numbers := make([]int,0)
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers,number)
	}
	numbers=append(numbers,0)
	sort.Ints(numbers)
	numbers=append(numbers,numbers[len(numbers)-1]+3)
	for i, number := range numbers {
		if i > 0 {
			if number-numbers[i-1] == 1 {
				one++
			}
			if number-numbers[i-1] == 3 {
				three++
			}

		}
	}
	return one*three
	panic("Not found")
}

var found = make(map[int]int)
func pathsToEnd(current int, paths map[int][]int) int {
	steps, wasFoundBefore := found[current]
	nextSteps := paths[current]
	if wasFoundBefore {
		return steps
	}

	if len(nextSteps) == 0 {
		return 1
 	}

	steps = 0
	for _, path := range nextSteps {
		steps += pathsToEnd(path,paths)
	}
	found[current] = steps

	return steps
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	pathMap := make(map[int][]int)

	lines := strings.Split(input, "\n")
	numbers := make([]int,0)


	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers,number)
	}

	numbers = append(numbers,0)
	sort.Ints(numbers)
	numbers=append(numbers,numbers[len(numbers)-1]+3)
	for i, number1 := range numbers {
		for j, number2 := range numbers {
			if i == j {
				continue
			}
			if number2 < number1 {
				continue
			}
			if number2 - number1 <= 3 {
				pathMap[number1] = append(pathMap[number1],number2)
			} else {
				break
			}

		}
	}
	return pathsToEnd(0,pathMap)
	panic("Not found")
}

func main() {
	start := time.Now()
	fmt.Println(filepath.Base(""))
	input := utils.Input(2020,10)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + strconv.Itoa(Part2(input)),"\t", time.Since(start))
}

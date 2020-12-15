package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type number struct {
	number      string
	first, last int
}

func parseNumbers(input string) (map[string]number, string, int) {
	numbers := make(map[string]number)
	lastNumber := ""
	turn := 1
	for _, s := range strings.Split(input, ",") {
		numbers[s] = number{first: turn, number: s, last: -1}
		turn++
		lastNumber = s
	}
	return numbers, lastNumber, turn
}

func incNumbers(numbers map[string]number, lastNumber string, turn int) (map[string]number, string) {
	n, ok := numbers[lastNumber]
	if !ok {
		numbers[lastNumber] = number{first: turn, last: -1, number: lastNumber}
		updateLastSaid("0", numbers, turn)
		return numbers, "0"
	} else {
		if n.last == -1 {
			updateLastSaid("0", numbers, turn)
			return numbers, "0"
		} else {
			nextNumber := strconv.Itoa(n.last - n.first)
			_, ok := numbers[nextNumber]
			if ok {
				updateLastSaid(nextNumber, numbers, turn)
			} else {
				numbers[nextNumber] = number{first: turn, last: -1, number: nextNumber}
			}

			return numbers, nextNumber
		}

	}
}

func updateLastSaid(number string, numbers map[string]number, turn int) {
	n := numbers[number]
	if n.last == -1 {
		n.last = turn
	} else {
		n.first, n.last = n.last, turn
	}
	numbers[number] = n
}

// Part1 Part 1 of puzzle
func Part1(input string) int {

	numbers, lastNumber, turn := parseNumbers(input)
	for turn <= 2020 {
		numbers, lastNumber = incNumbers(numbers, lastNumber, turn)
		turn++
	}
	l, _ := strconv.Atoi(lastNumber)
	return l
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	numbers, lastNumber, turn := parseNumbers(input)
	for turn <= 30000000 {
		numbers, lastNumber = incNumbers(numbers, lastNumber, turn)
		turn++
	}
	l, _ := strconv.Atoi(lastNumber)
	return l
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 15)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: ", Part1(input), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "\t", time.Since(start))
}

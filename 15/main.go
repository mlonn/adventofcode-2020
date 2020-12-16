package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type number struct {
	first, last int
	seen        bool
}

type game map[int]number

func parseNumbers(input string) (game, int) {
	numbers := make(game)
	var last int
	for i, s := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(s)
		numbers[n] = number{first: i + 1}
		last = n
	}
	return numbers, last
}

func (g game) next(previous int) int {
	n := g[previous]
	if !n.seen {
		return 0
	} else {
		return n.last - n.first
	}

}

func (g game) update(next int, turn int) {
	n, ok := g[next]
	if !ok {
		g[next] = number{first: turn}
	} else {
		if !n.seen {
			n.last = turn
			n.seen = true
		} else {
			n.first, n.last = n.last, turn
		}
		g[next] = n
	}
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	game, _ := parseNumbers(input)
	previous := -1
	turn := len(game) + 1
	for turn <= 2020 {
		next := game.next(previous)
		game.update(next, turn)
		previous = next
		turn++
	}
	return previous
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	game, previous := parseNumbers(input)
	turn := len(game) + 1
	for turn <= 30000000 {
		next := game.next(previous)
		game.update(next, turn)
		previous = next
		turn++
	}
	return previous
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

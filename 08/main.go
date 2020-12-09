package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)



// Part1 Part 1 of puzzle
func Part1(input string) string {
	acc := 0
	var visited = make(map[int]bool)
	lines := strings.Split(input, "\n")
	i := 0
	for i < len(lines) {
		instructions := strings.Split(lines[i]," ")
		instruction, n := instructions[0], instructions[1]
		number, _ := strconv.Atoi(n)
		if visited[i] {
			break
		} else {
			visited[i] = true
		}
		switch instruction {
		case "acc":
			acc += number
			i += 1
			break
		case "jmp":
			i += number
			break
		case "nop":
			i += 1
			break
		}
	}

	return strconv.Itoa(acc)
}

type Instruction struct {
	code   string
	number int
}

// Part2 Part2 of puzzle
func Part2(input string) string {


	lines := strings.Split(input, "\n")
	instructions := make([]Instruction,0)
	for i, _ := range lines {
		split := strings.Split(lines[i], " ")
		instruction, n := split[0], split[1]
		number, _ := strconv.Atoi(n)
		instructions = append(instructions, Instruction{code: instruction, number: number})
	}
	for index,instruction := range instructions {
		if instruction.code == "jmp" {
			instructions[index].code = "nop"
		} else if instruction.code == "nop" {
			instructions[index].code = "jmp"
		}
		i := 0
		acc := 0
		visited := make(map[int]bool)
		terminated := true
		for i < len(instructions) {
			ist := instructions[i]
			if visited[i] {
				terminated = false
				break
			} else {
				visited[i] = true
			}
			switch ist.code {
			case "acc":
				acc += ist.number
				i += 1
				break
			case "jmp":
				i += ist.number
				break
			case "nop":
				i += 1
				break
			}
		}
		if terminated {
			return strconv.Itoa(acc)
		}
		instructions[index] = instruction
	}
	return "Not found"
}

func main() {
	start := time.Now()
	input := utils.Input(2020,8)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Parse data: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + Part1(input), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(input),"\t", time.Since(start))
}

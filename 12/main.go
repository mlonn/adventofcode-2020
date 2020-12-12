package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)


type Instruction int
const (
	N = 0
	E = 90
	S = 180
	W = 270
)
// Part1 Part 1 of puzzle
func Part1(input string) int {
	lines := strings.Split(input,"\n")
	EW := 0
	NS := 0
	direction := E

	for _, line := range lines {
		instruction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])


		switch instruction {
		case 'L':
			direction =  (direction-magnitude+360)%360
			break
		case 'R':
			direction = (direction+magnitude)%360
			break
		case 'F':

			switch direction {
			case N:
				NS += magnitude
				break
			case W:
				EW += magnitude
				break
			case S:
				NS -= magnitude
				break
			case E:
				EW -= magnitude
				break
			}
			break
		case 'N':
			NS += magnitude
			break
		case 'W':
			EW -= magnitude
			break
		case 'S':
			NS -= magnitude
			break
		case 'E':
			EW += magnitude
			break

		}

	}
	return int(math.Abs(float64(NS)) + math.Abs(float64(EW)))
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	lines := strings.Split(input,"\n")
	EW := 0
	NS := 0
	WPEW := 10
	WPNS := 1
	for _, line := range lines {
		instruction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])
		switch instruction {
		case 'L':
			switch magnitude {
			case 90:
				WPEW, WPNS = -WPNS, WPEW
			case 180:
				WPEW, WPNS = -WPEW, -WPNS
			case 270:
				WPEW, WPNS = WPNS, -WPEW
			}
			break
		case 'R':
			switch magnitude {
			case 90:
				WPEW, WPNS = WPNS, -WPEW
			case 180:
				WPEW, WPNS = -WPEW, -WPNS
			case 270:
				WPEW, WPNS = -WPNS, WPEW
				break
			}
		case 'F':
			EW += magnitude * WPEW
			NS += magnitude * WPNS
			break
		case 'N':
			WPNS += magnitude
			break
		case 'W':
			WPEW -= magnitude
			break
		case 'S':
			WPNS -= magnitude
			break
		case 'E':
			WPEW += magnitude
			break
		}

	}
	return int(math.Abs(float64(NS)) + math.Abs(float64(EW)))
}

func main() {
	start := time.Now()
	input := utils.Input(2020,12)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + strconv.Itoa(Part2(input)),"\t", time.Since(start))
}

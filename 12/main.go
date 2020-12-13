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
	x := 0
	y := 0
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
				y += magnitude
				break
			case W:
				x += magnitude
				break
			case S:
				y -= magnitude
				break
			case E:
				x -= magnitude
				break
			}
			break
		case 'N':
			y += magnitude
			break
		case 'W':
			x -= magnitude
			break
		case 'S':
			y -= magnitude
			break
		case 'E':
			x += magnitude
			break
		}

	}
	return int(math.Abs(float64(y)) + math.Abs(float64(x)))
}


// Part2 Part2 of puzzle
func Part2(input string) int {
	lines := strings.Split(input,"\n")
	x := 0
	y := 0
	wpx := 10
	wpy := 1
	for _, line := range lines {
		instruction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])
		switch instruction {
		case 'L':
			turn:= 0
			for turn < (magnitude/90) {
				wpx, wpy = -wpy, wpx
				turn++
			}
			break
		case 'R':
			turn:= 0
			for turn < (magnitude/90) {
				wpx, wpy = wpy, -wpx
				turn++
			}
			break
		case 'F':
			x += magnitude * wpx
			y += magnitude * wpy
			break
		case 'N':
			wpy += magnitude
			break
		case 'W':
			wpx -= magnitude
			break
		case 'S':
			wpy -= magnitude
			break
		case 'E':
			wpx += magnitude
			break
		}

	}
	return int(math.Abs(float64(y)) + math.Abs(float64(x)))
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

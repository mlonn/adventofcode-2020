package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)


type Grid struct {
	g [][]string
	w int
	h int
}

type Direction int
const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)
func parseGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])
	grid := make([][]string, h)

	for i, line := range lines {
		grid[i] = make([]string, w)
		for j, position := range line {
			grid[i][j] = string(position)
		}
	}
	return Grid{g :grid, w: w, h: h}
}

func (grid Grid) Occupied(x,y int) bool {
	return grid.g[y][x] == "#"
}

func (grid Grid) Empty(x,y int) bool {
	return grid.g[y][x] == "L"
}

func (grid Grid) Occupied2(x,y int, direction Direction) int {
	// lower bound
	if x == -1 || y == -1 {
		return 0
	}
	// upper bound
	if x >= grid.w || y >= grid.h{
		return 0
	}
	if grid.g[y][x] == "." {
		switch direction {
			case N:
				return grid.Occupied2(x,y-1,direction)
			case NE:
				return grid.Occupied2(x+1,y-1,direction)
			case E:
				return grid.Occupied2(x+1,y,direction)
			case SE:
				return grid.Occupied2(x+1,y+1,direction)
			case S:
				return grid.Occupied2(x,y+1,direction)
			case SW:
				return grid.Occupied2(x-1,y+1,direction)
			case W:
				return grid.Occupied2(x-1,y,direction)
			case NW:
				return grid.Occupied2(x-1,y-1,direction)
		}
	}
	if grid.g[y][x] == "#" {
		return 1
	}
	return 0
}

func (grid Grid) NextPart2(x, y int) string{
	occupied := 0
	occupied += grid.Occupied2(x,y-1,N)
	occupied += grid.Occupied2(x+1,y-1,NE)
	occupied += grid.Occupied2(x+1,y,E)
	occupied += grid.Occupied2(x+1,y+1,SE)
	occupied += grid.Occupied2(x,y+1,S)
	occupied += grid.Occupied2(x-1,y+1,SW)
	occupied += grid.Occupied2(x-1,y,W)
	occupied += grid.Occupied2(x-1,y-1,NW)
	if occupied == 0 && grid.Empty(x,y) {
		return "#"
	} else if occupied >= 5 && grid.Occupied(x,y){
		return "L"
	}
	return grid.g[y][x]
}
func (grid Grid) NextPart1(x, y int) string{
	occupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Dont check our self
			if j == 0 && i == 0 {
				continue
			}
			// lower bound
			if x+j == -1 || y+i == -1 {
				continue
			}
			// upper bound
			if x+j >= grid.w || y+i >= grid.h{
				continue
			}
			if  grid.Occupied(x+j, y+i) {
				occupied++
			}
		}
	}
	if occupied == 0 && grid.Empty(x,y) {
		return "#"
	} else if occupied >= 4 && grid.Occupied(x,y){
		return "L"
	}
	return grid.g[y][x]
}

func (grid Grid) GetOccupied() int{
	occupied := 0
	for y, row := range grid.g {
		for x, _ := range row {
			if grid.Occupied(x,y){
				occupied++
			}
		}
	}
	return occupied
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	current := parseGrid(input)
	next := parseGrid(input)
	unstable := true
	step := 0
	for unstable {
		for y, row := range current.g {
			for x, _ := range row {
				next.g[y][x] = current.NextPart1(x, y)
			}
		}
		if current.GetOccupied() == next.GetOccupied(){
			unstable = false
		}
		current, next = next, current
		step ++
	}
	fmt.Println("Steps: ", step)
	return current.GetOccupied()
	panic("Not found")
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	current := parseGrid(input)
	next := parseGrid(input)
	unstable := true
	step := 0
	for unstable {
		for y, row := range current.g {
			for x, _ := range row {
				next.g[y][x] = current.NextPart2(x, y)
			}
		}
		current, next = next, current
		if current.GetOccupied() == next.GetOccupied(){
			unstable = false
		}
		step ++
	}
	fmt.Println("Steps: ", step)
	return current.GetOccupied()
	panic("Not found")
}

func main() {
	start := time.Now()
	input := utils.Input(2020,11)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + strconv.Itoa(Part2(input)),"\t", time.Since(start))
}

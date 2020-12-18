package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y, z, w int
}

type world map[Point]bool

var maxX, minX, maxY, minY int
var cycles = 6

func parseWorld(input string) world {
	lines := strings.Split(input, "\n")

	w := make(map[Point]bool)
	minX = math.MaxInt64
	minY = math.MaxInt64
	for y, line := range lines {
		for x, position := range line {
			if string(position) == "#" {
				w[Point{x: x, y: y}] = true
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
			}

		}
	}

	return w
}

func (w world) next(p Point) bool {
	active := 0
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if (dx != 0) || (dy != 0) || (dz != 0) || (dw != 0) {
						if w[Point{x: p.x + dx, y: p.y + dy, z: p.z + dz, w: p.w + dw}] {
							active++
						}
					}
				}
			}
		}
	}

	if !w[p] {
		if active == 3 {
			return true
		}
	}

	if w[p] {
		if active != 2 && active != 3 {
			return false
		}
	}
	return w[p]
}

func (w world) getActive() int {
	active := 0
	for _, a := range w {
		if a {
			active++
		}
	}
	return active
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	current := parseWorld(input)
	next := parseWorld(input)
	cycle := 0
	minZ := 0
	maxZ := 0
	for cycle < cycles {
		minX--
		minY--
		minZ--
		maxX++
		maxY++
		maxZ++

		for z := minZ; z <= maxZ; z++ {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					p := Point{x: x, y: y, z: z}
					next[p] = current.next(p)
				}
			}
		}

		current, next = next, current
		cycle++
	}
	return current.getActive()
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	current := parseWorld(input)
	next := parseWorld(input)
	cycle := 0
	minW := 0
	maxW := 0
	minZ := 0
	maxZ := 0
	for cycle < cycles {
		minX--
		minY--
		minZ--
		minW--
		maxX++
		maxY++
		maxZ++
		maxW++
		for w := minW; w <= maxW; w++ {
			for z := minZ; z <= maxZ; z++ {
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						p := Point{x: x, y: y, z: z, w: w}
						next[p] = current.next(p)
					}
				}
			}
		}
		current, next = next, current
		cycle++
	}
	return current.getActive()
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 17)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: "+strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: "+strconv.Itoa(Part2(input)), "\t", time.Since(start))
}

package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Grid struct {
	g  [][][]string
	xl int
	yl int
	zl int
}
type Point struct {
	x, y, z, w int
}
type world map[Point]bool

var zMin, xMin, yMin int
var maxX, minX, maxY, minY int
var cycles = 6

func parseGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])
	yl := h + cycles*2
	xl := w + cycles*2
	zl := 1 + cycles*2
	grid := make([][][]string, zl)
	zMin = -cycles
	yMin = -cycles - h/2
	xMin = -cycles - w/2
	for z := zMin; z <= zMin+cycles*2; z++ {
		grid[z-zMin] = make([][]string, yl)
		for y := yMin; y < yMin+yl; y++ {
			grid[z-zMin][y-yMin] = make([]string, xl)
			for x := range grid[z-zMin][y-yMin] {
				grid[z-zMin][y-yMin][x] = "."
			}
		}
	}

	for y, line := range lines {
		for x, position := range line {
			grid[cycles][y+cycles][x+cycles] = string(position)

		}
	}

	return Grid{g: grid, xl: xl, yl: yl, zl: zl}
}
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

func (grid Grid) get(x, y, z int) string {
	if z-zMin < 0 || y-yMin < 0 || x-xMin < 0 {
		return "."
	}
	if x-xMin >= grid.xl || y-yMin >= grid.yl || z-zMin >= grid.zl {
		return "."
	}
	return grid.g[z-zMin][y-yMin][x-xMin]
}

func (grid Grid) set(x, y, z int, s string) {
	grid.g[z-zMin][y-yMin][x-xMin] = s
}

func (grid Grid) isActive(x, y, z int) bool {
	return grid.get(x, y, z) == "#"
}

func (grid Grid) isInactive(x, y, z int) bool {
	return grid.get(x, y, z) == "."
}

func (grid Grid) next(x, y, z int) string {
	active := 0
	for dz := z - 1; dz <= z+1; dz++ {
		for dy := y - 1; dy <= y+1; dy++ {
			for dx := x - 1; dx <= x+1; dx++ {
				if (dx != x) || (dy != y) || (dz != z) {
					if grid.isActive(dx, dy, dz) {
						active++
					}
				}
			}
		}
	}

	if grid.isInactive(x, y, z) {
		if active == 3 {
			return "#"
		}
	}

	if grid.isActive(x, y, z) {
		if active != 2 && active != 3 {
			return "."
		}
	}
	return grid.get(x, y, z)
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

func (grid Grid) getActive() int {
	active := 0
	for z, slice := range grid.g {
		for y, row := range slice {
			for x := range row {
				if grid.g[z][y][x] == "#" {
					active++
				}
			}
		}
	}
	return active
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
	current := parseGrid(input)
	next := parseGrid(input)
	cycle := 0
	//minW := 0
	//maxW := 0
	fmt.Println(maxX, minX, maxY, minY)
	for cycle < cycles {
		for z := zMin; z < zMin+len(current.g); z++ {
			for y := yMin; y < yMin+len(current.g[z-zMin]); y++ {
				for x := xMin; x < xMin+len(current.g[z-zMin][y-yMin]); x++ {
					next.set(x, y, z, current.next(x, y, z))
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

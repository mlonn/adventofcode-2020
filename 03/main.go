package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)


func FindTrees(rows []string, right int, down int) int{
	trees:= 0
	position := 0
	for i, line := range rows {
		if i%down != 0 {
			continue
		}
		current := position % len(line)
		if line[current] == '#' {
			trees++
		}
		position+=right
	}
	return trees
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	s := strings.Split(input, "\n")
	return "Answer " + strconv.Itoa(FindTrees(s,3,1))
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	s := strings.Split(input, "\n")
	trees1 := FindTrees(s,1,1)
	trees3 := FindTrees(s,3,1)
	trees5 := FindTrees(s,5,1)
	trees7 := FindTrees(s,7,1)
	trees12 := FindTrees(s,1,2)
	product := trees1 * trees3 * trees5 * trees7 * trees12
	return "Answer " + strconv.Itoa(product)
}

func main() {
	start := time.Now()
	bytes, _ := ioutil.ReadFile("03/input.txt")
	fmt.Println("Read file in: ", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + Part1(string(bytes)), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(string(bytes)),"Time", time.Since(start))
}

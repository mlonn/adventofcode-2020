package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)


// Part1 Part 1 of puzzle
func Part1(input string) int {
	lines := strings.Split(input,"\n")
	mem := make(map[int]uint64)
	a := uint64(0)
	o := uint64(0)
	for _, line := range lines {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			andMask := ""
			orMask := ""
			for _, c := range split[1] {
				if c =='X' {
					andMask += "1"
					orMask += "0"
				} else {
					andMask += "0"
					orMask += string(c)
				}
			}
			a ,_ = strconv.ParseUint(andMask,2,64)
			o ,_ = strconv.ParseUint(orMask,2,64)
		} else {
			m := strings.Split(split[0],"[")
			address,_ := strconv.Atoi(strings.TrimSuffix(m[1],"]"))
			value, _ := strconv.ParseUint(split[1],10,64)
			value &= a
			value |= o
			mem[address] = value
		}
	}
	sum := uint64(0)
	for _, u := range mem {
		sum+=u
	}

	return int(sum)
}



// Part2 Part2 of puzzle
func Part2(input string) int {
	return -1
}

func main() {
	start := time.Now()
	input := utils.Input(2020,14)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + strconv.Itoa(Part2(input)),"\t", time.Since(start))
}

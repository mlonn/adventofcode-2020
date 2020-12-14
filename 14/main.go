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
	lines := strings.Split(input, "\n")
	mem := make(map[int]uint64)
	a := uint64(0)
	o := uint64(0)
	for _, line := range lines {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			andMask := ""
			orMask := ""
			for _, c := range split[1] {
				if c == 'X' {
					andMask += "1"
					orMask += "0"
				} else {
					andMask += "0"
					orMask += string(c)
				}
			}
			a, _ = strconv.ParseUint(andMask, 2, 64)
			o, _ = strconv.ParseUint(orMask, 2, 64)
		} else {
			m := strings.Split(split[0], "[")
			address, _ := strconv.Atoi(strings.TrimSuffix(m[1], "]"))
			value, _ := strconv.ParseUint(split[1], 10, 64)
			value &= a
			value |= o
			mem[address] = value
		}
	}
	sum := uint64(0)
	for _, u := range mem {
		sum += u
	}

	return int(sum)
}

func makeAddressCombinations(mask string) []uint64 {
	masks := make([]uint64, 0)
	a := []rune(mask)
	b := []rune(mask)
	for i, c := range []rune(mask) {
		if c == 'X' {
			a[i] = '1'
			b[i] = '0'
			if !strings.Contains(string(a), "X") {
				address, _ := strconv.ParseUint(string(a), 2, 64)
				masks = append(masks, address)
			} else {
				masks = append(masks, makeAddressCombinations(string(a))...)
			}
			if !strings.Contains(string(b), "X") {
				address, _ := strconv.ParseUint(string(b), 2, 64)
				masks = append(masks, address)
			} else {
				masks = append(masks, makeAddressCombinations(string(b))...)
			}
			break
		}

	}
	return masks
}

// Part2 Part2 of puzzle
func Part2(input string) uint64 {
	lines := strings.Split(input, "\n")
	mem := make(map[uint64]uint64)
	mask := []rune("")
	for _, line := range lines {
		split := strings.Split(line, " = ")
		value, _ := strconv.ParseUint(split[1], 10, 64)
		if split[0] == "mask" {
			mask = []rune(split[1])
		} else {
			m := strings.Split(split[0], "[")
			address, _ := strconv.ParseUint(strings.TrimSuffix(m[1], "]"), 10, 64)
			addressString := []rune(strconv.FormatUint(address, 2))
			addressString = []rune(fmt.Sprintf("%0*s", 36, string(addressString)))
			appliedMask := []rune(fmt.Sprintf("%0*s", 36, ""))
			for i := range addressString {
				if mask[i] == '0' {
					appliedMask[i] = addressString[i]

				} else {
					appliedMask[i] = mask[i]
				}
			}
			addressCombinations := makeAddressCombinations(string(appliedMask))
			for _, address := range addressCombinations {
				mem[address] = value
			}

		}
	}
	sum := uint64(0)
	for _, u := range mem {
		sum += u
	}

	return sum
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 14)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: ", Part1(input), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "\t", time.Since(start))
}

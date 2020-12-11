package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Bag struct {
	color string
	count int
	bags []Bag
}


func parseBags(input string) map[string]Bag{
	bags := make(map[string]Bag)
	bagRegex, _ := regexp.Compile("(\\d+) (.*) bags?")
	for _, bagString := range strings.Split(input, "\n") {
		split := strings.Split(bagString, " bags contain")
		color := split[0]
		contains := strings.Split(split[1], ", ")
		bag := Bag{color: color}
		for _, contain := range contains {
			match := bagRegex.FindStringSubmatch(contain)
			if match != nil {
				count, _ := strconv.Atoi(match[1])
				bag.bags = append(bag.bags, Bag{color: match[2], count: count})
			}

		}
		bags[bag.color] = bag
	}
	return bags
}

func canContainGold(color string, bags map[string]Bag) bool{
	if color == "shiny gold" {
		return true
	}
	for _, containedBag := range bags[color].bags {
		if canContainGold(containedBag.color, bags){
			return true
		}
	}
	return false
}

func countBags(color string, bags map[string]Bag) int{
	count := 0
	bag := bags[color]
	for _, containedBag := range bag.bags {
		subCount := countBags(containedBag.color, bags)
		count += containedBag.count + containedBag.count * subCount
	}
	return count
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	bags :=parseBags(input)
	sum := 0
	for _, bag := range bags {
		if canContainGold(bag.color, bags) && bag.color != "shiny gold" {
			sum++
		}
	}
	return sum
}



// Part2 Part2 of puzzle
func Part2(input string) int {
	bags := parseBags(input)
	return countBags("shiny gold", bags)
}

func main() {
	start := time.Now()
	input := utils.Input(2020,7)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: ", Part1(input), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input),"\t", time.Since(start))
}

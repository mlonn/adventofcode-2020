package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
func parseBags(input string) map[string]Bag {
	bagRegex, _ := regexp.Compile("(\\d+) (.*) bags?")
	bags := make(map[string]Bag)
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

func canContainGold(bagMap map[string]Bag, color string) bool{
	if color == "shiny gold" {
		return true
	}
	for _, containedBag := range bagMap[color].bags {
		if canContainGold(bagMap, containedBag.color){
			return true
		}
	}
	return false
}

func countBags(bagMap map[string]Bag, color string) int{
	count := 0
	bag := bagMap[color]
	for _, containedBag := range bag.bags {
		subCount := countBags(bagMap, containedBag.color)
		count += containedBag.count + containedBag.count * subCount
	}
	return count
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	sum := 0
	bagMap := parseBags(input)
	for _, bag := range bagMap {
		if canContainGold(bagMap, bag.color) && bag.color != "shiny gold" {
			sum++
		}
	}
	return "Answer " + strconv.Itoa(sum)
}



// Part2 Part2 of puzzle
func Part2(input string) string {
	bags := parseBags(input)
	return "Answer " + strconv.Itoa(countBags(bags, "shiny gold"))
}

func main() {
	start := time.Now()
	data, _ := ioutil.ReadFile("07/input.txt")
	input := string(bytes.TrimSpace(data))
	fmt.Println("Read file in: ", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(input),"Time", time.Since(start))
}

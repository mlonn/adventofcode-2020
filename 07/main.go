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
var bags = make(map[string]Bag)

func parseBags(input string) {
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
}

func canContainGold(color string) bool{
	if color == "shiny gold" {
		return true
	}
	for _, containedBag := range bags[color].bags {
		if canContainGold(containedBag.color){
			return true
		}
	}
	return false
}

func countBags(color string) int{
	count := 0
	bag := bags[color]
	for _, containedBag := range bag.bags {
		subCount := countBags(containedBag.color)
		count += containedBag.count + containedBag.count * subCount
	}
	return count
}

// Part1 Part 1 of puzzle
func Part1() string {
	sum := 0
	for _, bag := range bags {
		if canContainGold(bag.color) && bag.color != "shiny gold" {
			sum++
		}
	}
	return "Answer " + strconv.Itoa(sum)
}



// Part2 Part2 of puzzle
func Part2() string {
	return "Answer " + strconv.Itoa(countBags("shiny gold"))
}

func main() {
	start := time.Now()
	data, _ := ioutil.ReadFile("07/input.txt")
	input := string(bytes.TrimSpace(data))
	fmt.Println("Read file: ", time.Since(start))
	start = time.Now()
	parseBags(input)
	fmt.Println("Parsed data: ", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: " + Part1(), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(),"Time", time.Since(start))
}

package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type span struct {
	min, max int
}

type rule struct {
	lower, upper span
}
type rules map[string]rule

type ticket []int

type possibilities map[int][]string

func parseTicket(ticketString string) ticket {
	ticketSplit := strings.Split(ticketString, ",")
	returnTicket := make(ticket, len(ticketSplit))
	for i, s := range ticketSplit {
		number, _ := strconv.Atoi(s)
		returnTicket[i] = number
	}
	return returnTicket
}

func parseSpan(s string) span {
	minmax := strings.Split(s, "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])
	return span{min: min, max: max}
}

func parse(input string) (rules, ticket, []ticket) {
	parts := strings.Split(input, "\n\n")
	rulePart := strings.Split(parts[0], "\n")
	rules := make(rules)
	for _, part := range rulePart {
		split := strings.Split(part, ": ")
		key := split[0]
		ruleSplit := strings.Split(split[1], " or ")
		lower := parseSpan(ruleSplit[0])
		upper := parseSpan(ruleSplit[1])
		rules[key] = rule{lower: lower, upper: upper}
	}
	myTicket := parseTicket(strings.Split(parts[1], "\n")[1])
	nearByTicketsSplit := strings.Split(parts[2], "\n")[1:]
	nearByTickets := make([]ticket, len(nearByTicketsSplit))
	for i, nt := range nearByTicketsSplit {
		nearByTickets[i] = parseTicket(nt)
	}
	return rules, myTicket, nearByTickets
}

func (r span) isInSpan(n int) bool {
	return r.min <= n && n <= r.max
}

func (t ticket) getErrorSum(r rules) int {
	errorSum := 0
	for _, n := range t {
		valid := false
		for _, rule := range r {
			if rule.lower.isInSpan(n) || rule.upper.isInSpan(n) {
				valid = true
			}
		}
		if !valid {
			errorSum += n
		}
	}
	return errorSum
}

func (t ticket) getValidRules(r rules) map[int][]string {
	validForPositions := make(map[int]map[string]bool)
	for i, n := range t {
		validRules := make(map[string]bool)
		for name, rule := range r {
			if rule.lower.isInSpan(n) || rule.upper.isInSpan(n) {
				validRules[name] = true
			}
		}
		if len(validRules) != 0 {
			validForPositions[i] = validRules
		}
	}
	valid := make(map[int][]string)
	for i, set := range validForPositions {
		list := make([]string, 0)
		for s := range set {
			list = append(list, s)
		}
		valid[i] = list
	}
	return valid
}

func (p possibilities) removeFound(i int, found string) {
	for j, keys := range p {
		if i != j {
			newKeys := make([]string, 0)
			for _, i := range keys {
				if i != found {
					newKeys = append(newKeys, i)
				}
			}
			p[j] = newKeys
		}

	}
}

func (p possibilities) update(i int, val []string, found int) int {
	if rules, exists := p[i]; exists {
		intersect := intersection(rules, val)
		p[i] = intersect
		if len(intersect) == 1 {
			found++
			p.removeFound(i, intersect[0])
		}
	} else {
		p[i] = val
	}
	return found
}

func intersection(a []string, b []string) []string {
	set := make(map[string]bool)
	hash := make(map[string]bool)
	for _, el := range a {
		hash[el] = true
	}
	for _, el := range b {
		if _, found := hash[el]; found {
			set[el] = true
		}
	}
	list := make([]string, 0)
	for s := range set {
		list = append(list, s)
	}
	return list
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	rules, _, nearByTickets := parse(input)
	errorSum := 0
	for _, t := range nearByTickets {
		errorSum += t.getErrorSum(rules)
	}
	return errorSum
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	r, myTicket, nearByTickets := parse(input)
	possible := make(possibilities)
	found := 0
	for found <= len(r) {
		for _, t := range nearByTickets {
			validRules := t.getValidRules(r)
			for i, val := range validRules {
				found = possible.update(i, val, found)
			}
		}
	}
	sum := 1
	for position, rule := range possible {
		if strings.Contains(rule[0], "departure") {
			sum *= myTicket[position]
		}
	}
	return sum
}

func main() {
	start := time.Now()
	input := utils.Input(2020, 16)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: ", Part1(input), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "\t", time.Since(start))
}

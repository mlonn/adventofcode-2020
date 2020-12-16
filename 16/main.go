package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"reflect"
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

func makeTicket(ticketString string) ticket {
	ticketSplit := strings.Split(ticketString, ",")
	returnTicket := make(ticket, len(ticketSplit))
	for i, s := range ticketSplit {
		number, _ := strconv.Atoi(s)
		returnTicket[i] = number
	}
	return returnTicket
}

func makeSpan(s string) span {
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
		lower := makeSpan(ruleSplit[0])
		upper := makeSpan(ruleSplit[1])
		rules[key] = rule{lower: lower, upper: upper}
	}

	myTicket := makeTicket(strings.Split(parts[1], "\n")[1])
	nearByTicketsSplit := strings.Split(parts[2], "\n")[1:]
	nearByTickets := make([]ticket, len(nearByTicketsSplit))
	for i, nt := range nearByTicketsSplit {
		nearByTickets[i] = makeTicket(nt)
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
func (t ticket) getValidRules(r rules) map[int]map[string]bool {
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
	return validForPositions
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

func intersection(a map[string]bool, b map[string]bool) map[string]bool {
	set := make(map[string]bool)
	hash := make(map[string]bool)
	for el := range a {
		hash[el] = true
	}
	for el := range b {
		if _, found := hash[el]; found {
			set[el] = true
		}
	}
	return set
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	r, myTicket, nearByTickets := parse(input)
	valid := make(map[int]map[string]bool)
	positions := make(map[int]string)
	for len(positions) < len(r) {
		for _, t := range nearByTickets {
			ticketPosition := t.getValidRules(r)
			for i, val := range ticketPosition {
				if _, found := valid[i]; found {
					intersect := intersection(valid[i], val)
					valid[i] = intersect
					if len(intersect) == 1 {
						key := reflect.ValueOf(intersect).MapKeys()[0].String()
						positions[i] = key
						for j, v := range valid {
							delete(v, key)
							valid[j] = v
						}
					}
				} else {
					valid[i] = val
				}
			}
		}
	}
	sum := 1
	for position, rule := range positions {
		if strings.Contains(rule, "departure") {
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

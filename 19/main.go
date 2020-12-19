package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type exp struct {
	op          string
	left, right string
	expressions []exp
}

// Part1 Part 1 of puzzle
func Part1(input string) int {
	lines := strings.Split(input, "\n")
	expressions := make([]exp, len(lines))
	for i, line := range lines {
		e := exp{}
		e.parse(line)
		expressions[i] = e
	}
	for _, exp := range expressions {
		exp.eval()
	}
	return -1
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	return -1
}
func (e exp) eval() int {
	for _, ee := range e.expressions {
		fmt.Println(ee.eval())
	}
	fmt.Println(e.left)
	return -1
}
func (e *exp) parse(line string) {
	line = strings.ReplaceAll(line, " ", "")
	s := ""
	i := 0
	for i < len(line) {
		c := line[i]
		if c == '(' {
			sub := splitParen(line[i:])
			e.append(sub)
			i += len(sub) + 2
		} else if c == '*' || c == '+' {
			e.op = string(c)
			e.left = s
			s = ""
			i++
		} else {
			s += string(c)
			i++
		}
	}
	if s != "" {
		e.right = s
	}
}

func (e *exp) append(s string) {
	if strings.Contains(s, "(") || strings.Contains(s, "*") || strings.Contains(s, "+") {
		ec := exp{}
		ec.parse(s)
		e.expressions = append(e.expressions, ec)
	} else {
		ec := exp{left: s}
		e.expressions = append(e.expressions, ec)
	}
}
func splitParen(line string) string {
	openParen := 0
	for i, c := range line {
		if c == ')' {
			openParen++
		}
		if c == '(' {
			openParen--
		}
		if openParen == 0 {
			return line[1:i]
		}

	}
	return line
}
func main() {
	start := time.Now()
	input := utils.Input(2020, 18)
	fmt.Println("Read file: \t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1: "+strconv.Itoa(Part1(input)), "\t", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: "+strconv.Itoa(Part2(input)), "\t", time.Since(start))
}

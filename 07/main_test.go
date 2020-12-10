package main

import (
	"advent-of-code-2020/utils"
	"testing"
)


func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/7.txt")
	solution := 302
	parseBags(input)
	for n := 0; n < b.N; n++ {
		answer := Part1()
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/7.txt")
	solution := 4165
	parseBags(input)
	for n := 0; n < b.N; n++ {
		answer := Part2()
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}
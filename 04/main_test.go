package main

import (
	"advent-of-code-2020/utils"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/4.txt")
	solution := 219
	for n := 0; n < b.N; n++ {
		answer := Part1(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/4.txt")
	solution := 127
	for n := 0; n < b.N; n++ {
		answer := Part2(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

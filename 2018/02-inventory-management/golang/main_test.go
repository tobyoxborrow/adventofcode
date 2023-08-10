package aoc2018_02

import (
	"fmt"
	"testing"
)

var testCasesA = []struct {
	input    []string
	expected int
}{
	{[]string{
		"abcdef", // contains no letters that appear exactly two or three times.
		"bababc", // contains two a and three b, so it counts for both.
		"abbcde", // contains two b, but no letter appears exactly three times.
		"abcccd", // contains three c, but no letter appears exactly two times.
		"aabcdd", // contains two a and two d, but it only counts once.
		"abcdee", // contains two e.
		"ababab", // contains three a and three b, but it only counts once.
	}, 12},
}

func TestSolveA(t *testing.T) {
	for _, tc := range testCasesA {
		name := fmt.Sprintf("%v", tc.input[0])
		t.Run(name, func(t *testing.T) {
			result := SolveA(tc.input)
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkSolveA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveA(testCasesA[0].input)
	}
}

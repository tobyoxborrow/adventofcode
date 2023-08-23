package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveATable(t *testing.T) {
	cases := []struct {
		polymer  string
		expected int
	}{
		{"aA", 0},
		{"abBA", 0},
		{"aabBCC", 4},
		{"abAB", 4},
		{"aabAAB", 6},
		{"aabbccddeeffgghhiijjJJIIHHGGFFEEDDCCBBAA", 0},
		{"aabbccdaabbccddeeffgghhiijjJJIIHHGGFFEEDDCCBBAAdeeffgghhiijjJJIIHHGGFFEEDDCCBBAA", 0},
	}

	for _, tc := range cases {
		name := string(tc.polymer)
		t.Run(name, func(t *testing.T) {
			result := SolveA([]byte(tc.polymer))
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSolveA(t *testing.T) {
	expected := 10

	lines := parseInput(sample)

	result := SolveA(lines)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveBTable(t *testing.T) {
	cases := []struct {
		polymer  string
		expected int
	}{
		{"aA", 0},
		{"abBA", 0},
		{"AzzAZz", 2},
		{"ajabbccddeeffgghhiijjzJJIIHHGGFFEEDDCCBBAAZJ", 4},
	}

	for _, tc := range cases {
		name := string(tc.polymer)
		t.Run(name, func(t *testing.T) {
			result := SolveB([]byte(tc.polymer))
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSolveB(t *testing.T) {
	expected := 4

	lines := parseInput(sample)

	result := SolveB(lines)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInput(sample)
	}
}

func BenchmarkSolveA(b *testing.B) {
	lines := parseInput(sample)
	for i := 0; i < b.N; i++ {
		SolveA(lines)
	}
}

func BenchmarkSolveB(b *testing.B) {
	lines := parseInput(sample)
	for i := 0; i < b.N; i++ {
		SolveB(lines)
	}
}

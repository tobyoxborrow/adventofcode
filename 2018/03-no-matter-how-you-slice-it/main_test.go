package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	expected := 4

	lines := parseInput(sample)
	claims := makeClaimsFromLines(lines)
	grid := drawGrid(claims)

	result := SolveA(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 3

	lines := parseInput(sample)
	claims := makeClaimsFromLines(lines)
	grid := drawGrid(claims)

	result := SolveB(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInput(sample)
	}
}

func BenchmarkMakeClaims(b *testing.B) {
	lines := parseInput(sample)
	for i := 0; i < b.N; i++ {
		makeClaimsFromLines(lines)
	}
}

func BenchmarkDrawGrid(b *testing.B) {
	lines := parseInput(sample)
	claims := makeClaimsFromLines(lines)
	for i := 0; i < b.N; i++ {
		drawGrid(claims)
	}
}

func BenchmarkSolveA(b *testing.B) {
	lines := parseInput(sample)
	claims := makeClaimsFromLines(lines)
	grid := drawGrid(claims)
	for i := 0; i < b.N; i++ {
		SolveA(grid)
	}
}

func BenchmarkSolveB(b *testing.B) {
	lines := parseInput(sample)
	claims := makeClaimsFromLines(lines)
	grid := drawGrid(claims)
	for i := 0; i < b.N; i++ {
		SolveB(grid)
	}
}

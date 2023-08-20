package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	expected := 240

	lines := parseInput(sample)
	guards := makeGuards(lines)

	result := SolveA(guards)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 4455

	lines := parseInput(sample)
	guards := makeGuards(lines)

	result := SolveB(guards)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInput(sample)
	}
}

func BenchmarkMakeGuards(b *testing.B) {
	lines := parseInput(sample)
	for i := 0; i < b.N; i++ {
		makeGuards(lines)
	}
}

func BenchmarkSolveA(b *testing.B) {
	lines := parseInput(sample)
	guards := makeGuards(lines)
	for i := 0; i < b.N; i++ {
		SolveA(guards)
	}
}

func BenchmarkSolveB(b *testing.B) {
	lines := parseInput(sample)
	guards := makeGuards(lines)
	for i := 0; i < b.N; i++ {
		SolveB(guards)
	}
}

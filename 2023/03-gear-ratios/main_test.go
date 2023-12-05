package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveOne(t *testing.T) {
	expected := 4361

	grid := parseInput(sample)

	result := SolveOne(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveTwo(t *testing.T) {
	expected := 467835

	grid := parseInput(sample)

	result := SolveTwo(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

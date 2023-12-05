package main

import (
	_ "embed"
	"testing"
)

//go:embed sampled.txt
var sample string

func TestSolveOne(t *testing.T) {
	expected := 4361

	grid := parseInput(sample)

	result := SolveOne(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

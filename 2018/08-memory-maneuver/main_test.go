package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	expected := 138

	numbers := parseInput(sample)

	result := SolveA(numbers)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 66

	numbers := parseInput(sample)

	result := SolveB(numbers)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

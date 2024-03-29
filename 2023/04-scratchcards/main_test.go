package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveOne(t *testing.T) {
	expected := 13

	parsedInput := parseInput(sample)

	result := SolveOne(parsedInput)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveTwo(t *testing.T) {
	expected := 30

	parsedInput := parseInput(sample)

	result := SolveTwo(parsedInput)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

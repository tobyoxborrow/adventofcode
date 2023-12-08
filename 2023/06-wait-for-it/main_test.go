package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveOne(t *testing.T) {
	expected := 288

	parsedInput := parseInputOne(sample)

	result := Solve(parsedInput)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveTwo(t *testing.T) {
	expected := 71503

	parsedInput := parseInputTwo(sample)

	result := Solve(parsedInput)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

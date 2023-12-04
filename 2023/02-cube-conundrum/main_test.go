package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	expected := 8

	games := parseInput(sample)

	result := SolveA(games)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 2286

	games := parseInput(sample)

	result := SolveB(games)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

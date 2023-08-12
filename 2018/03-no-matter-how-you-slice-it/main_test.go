package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	expected := 4
	result := SolveA(sample)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 3
	result := SolveB(sample)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

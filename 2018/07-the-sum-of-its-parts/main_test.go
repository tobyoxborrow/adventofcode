package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

/*
func TestParseLocations(t *testing.T) {
	cases := []struct {
		line      []string
		expectedX int
		expectedY int
	}{
		{[]string{"1, 2\n"}, 1, 2},
		{[]string{"1, 2"}, 1, 2},
		{[]string{"10, 20"}, 10, 20},
		{[]string{"123, 456"}, 123, 456},
		{[]string{"1234, 20"}, 1234, 20},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%s", tc.line)
		t.Run(name, func(t *testing.T) {
			locations := newLocations(tc.line)
			if len(locations) != 1 {
				t.Fatalf("Expected 1 location but got %d\n", len(locations))
			}
			result := locations[0]
			if result.point.x != tc.expectedX {
				t.Fatalf("expected X %v, got %v", tc.expectedX, result)
			}
			if result.point.y != tc.expectedY {
				t.Fatalf("expected Y %v, got %v", tc.expectedY, result)
			}
		})
	}
}
*/

func TestSolveA(t *testing.T) {
	expected := "CABDFE"

	lines := parseInput(sample)
	list := newAdjacencyList(lines)

	result := SolveA(list)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 15

	lines := parseInput(sample)
	list := newAdjacencyList(lines)

	result := SolveB(list, 2, 0)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func BenchmarkSolveA(b *testing.B) {
	lines := parseInput(sample)
	list := newAdjacencyList(lines)
	for i := 0; i < b.N; i++ {
		SolveA(list)
	}
}

func BenchmarkSolveB(b *testing.B) {
	lines := parseInput(sample)
	list := newAdjacencyList(lines)
	for i := 0; i < b.N; i++ {
		SolveB(list, 5, 60)
	}
}

package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed sample.txt
var sample string

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

func TestDistanceFrom(t *testing.T) {
	cases := []struct {
		x1       int
		y1       int
		x2       int
		y2       int
		expected int
	}{
		{1, 1, 0, 0, 2},
		{1, 1, 1, 0, 1},
		{1, 1, 2, 0, 2},
		{1, 1, 3, 0, 3},
		{1, 1, 4, 0, 4},
		{1, 1, 5, 0, 5},
		{8, 3, 0, 0, 11},
		{8, 3, 1, 0, 10},
		{8, 3, 2, 0, 9},
		{8, 3, 3, 0, 8},
		{8, 3, 4, 0, 7},
		{8, 3, 5, 0, 6},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%dx%d-%dx%d", tc.x1, tc.y1, tc.x2, tc.y2)
		t.Run(name, func(t *testing.T) {
			point1 := Point{tc.x1, tc.y1}
			location := Location{point: point1}
			point2 := Point{tc.x2, tc.y2}
			result := location.DistanceFrom(point2)
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSolveA(t *testing.T) {
	expected := 17

	lines := parseInput(sample)
	locations := newLocations(lines)
	grid := newGrid(locations)
	grid.populateDistances()

	result := SolveA(grid)
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSolveB(t *testing.T) {
	expected := 16

	lines := parseInput(sample)
	locations := newLocations(lines)
	grid := newGrid(locations)
	grid.populateDistances()

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

func BenchmarkNewLocations(b *testing.B) {
	lines := parseInput(sample)
	for i := 0; i < b.N; i++ {
		newLocations(lines)
	}
}

func BenchmarkPopulateDistances(b *testing.B) {
	lines := parseInput(sample)
	locations := newLocations(lines)
	grid := newGrid(locations)
	for i := 0; i < b.N; i++ {
		grid.populateDistances()
	}
}

func BenchmarkSolveA(b *testing.B) {
	lines := parseInput(sample)
	locations := newLocations(lines)
	grid := newGrid(locations)
	grid.populateDistances()
	for i := 0; i < b.N; i++ {
		SolveA(grid)
	}
}

func BenchmarkSolveB(b *testing.B) {
	lines := parseInput(sample)
	locations := newLocations(lines)
	grid := newGrid(locations)
	grid.populateDistances()
	for i := 0; i < b.N; i++ {
		SolveB(grid)
	}
}

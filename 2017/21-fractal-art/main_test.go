package main

import "testing"

var sample1 = []string{
	"../.# => ##./#../...",
	".#./..#/### => #..#/..../..../#..#",
}

func TestSolve(t *testing.T) {
	if solve(sample1, 2) != 12 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1, 2)
	}
}

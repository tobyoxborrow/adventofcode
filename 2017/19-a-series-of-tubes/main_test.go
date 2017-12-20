package main

import "testing"

var sample1 = []string{
	"     |",
	"     |  +--+",
	"     A  |  C",
	" F---|----E|--+",
	"     |  |  |  D",
	"     +B-+  +--+",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != "ABCDEF" {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1)
	}
}

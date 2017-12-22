package main

import "testing"

var sample1 = []string{
	"..#",
	"#..",
	"...",
}

func TestSolveA1(t *testing.T) {
	if solve(sample1, 7) != 5 {
		t.Fail()
	}
}

func TestSolveA2(t *testing.T) {
	if solve(sample1, 70) != 41 {
		t.Fail()
	}
}

func TestSolveA3(t *testing.T) {
	if solve(sample1, 10000) != 5587 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1, 10000)
	}
}

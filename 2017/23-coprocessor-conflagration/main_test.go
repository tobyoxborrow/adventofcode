package main

import "testing"

var sample1 = []string{
	"set b 1",
	"mul b 2",
	"mul b 2",
	"set h 8",
}

func TestSolve(t *testing.T) {
	if solve(sample1, debugOff) != 2 {
		t.Fail()
	}
}

func TestSolveB(t *testing.T) {
	if solve(sample1, debugOn) != 8 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1, debugOff)
	}
}

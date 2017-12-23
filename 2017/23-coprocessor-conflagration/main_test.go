package main

import "testing"

var sample1 = []string{
	"set a 1",
	"mul a 2",
	"mul a 2",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != 2 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1)
	}
}

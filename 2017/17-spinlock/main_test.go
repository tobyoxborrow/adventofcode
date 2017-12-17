package main

import "testing"

func TestSolve(t *testing.T) {
	if solve(3) != 638 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(3)
	}
}

func BenchmarkSolveB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveB(3)
	}
}

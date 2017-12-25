package main

import "testing"

var sample1 = []string{
	"0/2",
	"2/2",
	"2/3",
	"3/4",
	"3/5",
	"0/1",
	"10/1",
	"9/10",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != 31 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(getChallenge())
	}
}

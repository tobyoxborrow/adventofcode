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
	solve(sample1)
	if answerA != 31 {
		t.Fail()
	}
}

func TestSolveB(t *testing.T) {
	solve(sample1)
	if answerB != 19 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(getChallenge())
	}
}

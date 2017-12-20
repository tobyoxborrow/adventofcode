package main

import "testing"

var sample1 = []string{
	"set a 1",
	"add a 2",
	"mul a a",
	"mod a 5",
	"snd a",
	"set a 0",
	"rcv a",
	"jgz a -1",
	"set a 1",
	"jgz a -2",
}

var sample2 = []string{
	"snd 1",
	"snd 2",
	"snd p",
	"rcv a",
	"rcv b",
	"rcv c",
	"rcv d",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != 4 {
		t.Fail()
	}
}

func TestSolveB(t *testing.T) {
	if solveB(sample2) != 3 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1)
	}
}

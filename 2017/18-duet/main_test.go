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

func TestSolve(t *testing.T) {
	if solve(sample1) != 4 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1)
	}
}

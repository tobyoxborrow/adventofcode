package main

import "testing"

var sample1 = []string{
	"p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
	"p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != 0 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(sample1)
	}
}

package main

import "testing"

func TestSolve(t *testing.T) {
	if solve("1122") != 3 {
		t.Fail()
	}
}

func TestSolve2(t *testing.T) {
	if solve("1111") != 4 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve("1122")
	}
}

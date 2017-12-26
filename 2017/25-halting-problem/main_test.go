package main

import "testing"

var sample1 = []string{
	"Begin in state A.",
	"Perform a diagnostic checksum after 6 steps.",
	"",
	"In state A:",
	"  If the current value is 0:",
	"    - Write the value 1.",
	"    - Move one slot to the right.",
	"    - Continue with state B.",
	"  If the current value is 1:",
	"    - Write the value 0.",
	"    - Move one slot to the left.",
	"    - Continue with state B.",
	"",
	"In state B:",
	"  If the current value is 0:",
	"    - Write the value 1.",
	"    - Move one slot to the left.",
	"    - Continue with state A.",
	"  If the current value is 1:",
	"    - Write the value 1.",
	"    - Move one slot to the right.",
	"    - Continue with state A.",
	"",
}

func TestSolve(t *testing.T) {
	if solve(sample1) != 3 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(getChallenge())
	}
}

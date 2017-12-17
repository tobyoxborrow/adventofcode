package main

import "testing"

func TestSolve(t *testing.T) {
	if solve(3) != 638 {
		t.Fail()
	}
}

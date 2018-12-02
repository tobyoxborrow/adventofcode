package main

import "testing"

var sample1 = []string{
	"+1",
	"+1",
	"+1",
}

var sample2 = []string{
	"+1",
	"+1",
	"-2",
}

var sample3 = []string{
	"-1",
	"-2",
	"-3",
}

var sample4 = []string{
	"+1",
	"-1",
}

var sample5 = []string{
	"+3",
	"+3",
	"+4",
	"-2",
	"-4",
}

var sample6 = []string{
	"-6",
	"+3",
	"+8",
	"+5",
	"-6",
}

func TestSolve1(t *testing.T) {
	if solve(sample1) != 3 {
		t.Fail()
	}
}

func TestSolve2(t *testing.T) {
	if solve(sample2) != 0 {
		t.Fail()
	}
}

func TestSolve3(t *testing.T) {
	if solve(sample3) != -6 {
		t.Fail()
	}
}

func TestSolve4(t *testing.T) {
	if solveB(sample4) != 0 {
		t.Fail()
	}
}

func TestSolve5(t *testing.T) {
	if solveB(sample5) != 10 {
		t.Fail()
	}
}

func TestSolve6(t *testing.T) {
	if solveB(sample6) != 5 {
		t.Fail()
	}
}

package main

import (
	"fmt"
	"testing"
)

func TestSolveA(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{`
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`, 12}}

	for _, tc := range cases {
		name := fmt.Sprintf("%v", tc.input[0])
		t.Run(name, func(t *testing.T) {
			result := SolveA(tc.input)
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{`
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
`, "fgij"}}

	for _, tc := range cases {
		name := fmt.Sprintf("%v", tc.input[0])
		t.Run(name, func(t *testing.T) {
			result := SolveB(tc.input)
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

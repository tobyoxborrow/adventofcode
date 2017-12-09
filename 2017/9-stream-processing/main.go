package main

/*

Day 9: Stream Processing

Group examples
{}, 1 group.
{{{}}}, 3 groups.

Group scoring
{}, score of 1.
{{{}}}, score of 1 + 2 + 3 = 6.

Garbage examples
<random characters>, garbage containing random characters.
<<<<>, because the extra < are ignored.
<{!>}>, because the first > is canceled.
<!!>, because the second ! is canceled, allowing the > to terminate the garbage

A:
What is the total score for all groups in your input?

B:
-

*/

import (
	"fmt"
	"io/ioutil"
)

func getChallenge() []byte {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func solve(stream []byte) (score int) {
	gr := 0     // group depth
	gb := false // inside garbage
	ig := false // ignore next char
	for _, b := range stream {
		if ig {
			ig = false
			continue
		}

		switch string(b) {
		case "{":
			if !gb {
				gr++
			}
		case "}":
			if !gb {
				score += gr
				gr--
			}
		case "<":
			gb = true
		case ">":
			gb = false
		case "!":
			ig = true
		}
	}
	return
}

func main() {
	testCase1 := []byte("{{{}}}")
	testCase2 := []byte("{{{},{},{{}}}}")
	testCase3 := []byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	testCase4 := []byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	challengeInput := getChallenge()

	fmt.Println(solve(testCase1) == 6)
	fmt.Println(solve(testCase2) == 16)
	fmt.Println(solve(testCase3) == 9)
	fmt.Println(solve(testCase4) == 3)
	fmt.Println(solve(challengeInput))
}

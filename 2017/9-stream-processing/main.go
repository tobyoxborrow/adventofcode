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
How many non-canceled characters are within the garbage in your puzzle input?

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

func solve(stream []byte) (score int, ncc int) {
	gr := 0     // group depth
	gb := false // inside garbage
	ig := false // ignore next char
	for _, b := range stream {
		if ig {
			ig = false
			continue
		}

		if gb {
			ncc++
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
			ncc--
		case "!":
			ig = true
			ncc--
		}
	}
	return
}

func main() {
	testCase1 := []byte("{{{}}}")
	testCase2 := []byte("{{{},{},{{}}}}")
	testCase3 := []byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	testCase4 := []byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	testCase5 := []byte("<<<<>")
	testCase6 := []byte("<{o'i!a,<{i<a>")
	challengeInput := getChallenge()

	fmt.Println(solve(testCase1)) // 6 0
	fmt.Println(solve(testCase2)) // 16 0
	fmt.Println(solve(testCase3)) // 9 8
	fmt.Println(solve(testCase4)) // 3 17
	fmt.Println(solve(testCase5)) // 0 3
	fmt.Println(solve(testCase6)) // 0 10
	fmt.Println(solve(challengeInput))
}

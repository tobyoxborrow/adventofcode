package main

/*

Day 13: Packet Scanners

0: 3
1: 2
4: 4
6: 4

You need to cross a vast firewall. The firewall consists of several layers,
each with a security scanner that moves back and forth across the layer. To
succeed, you must not be detected by a scanner.

A:
What is the severity of your whole trip?

B:
-

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	return strings.Split(s, "\n")
}

type layer struct {
	id        int // "range", but that is a reserved word
	depth     int
	severity  int  // range * depth
	scanner   int  // scanner position within the layer
	direction int8 // scanner current movement direction; 0 down, 1 up
}

func solve(lines []string) (score int) {
	// populate basic program lookup table from input lines
	lastID := 0
	layers := make(map[int]*layer)
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			panic("Bad range/ID in line")
		}
		depth, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("Bad depth in line")
		}
		// fill in any gaps, lucky the input is ordered
		for c := lastID + 1; c < id; c++ {
			layers[c] = &layer{c, 0, 0, -1, 0}
		}
		layers[id] = &layer{id, depth, id * depth, 0, 0}
		lastID = id
	}

	score = 0
	/*
		for p := 0; p <= lastID; p++ {
			if layers[p].depth == 0 {
				continue
			}
			fmt.Println(p, layers[p].depth)
		}
	*/

	// "packet" movement from left to right
	for p := 0; p <= lastID; p++ {
		// scoring
		if layers[p].scanner == 0 {
			score += layers[p].severity
			//fmt.Println("DING:", p, score)
		}

		// scanner activity
		for s := 0; s <= lastID; s++ {
			// skip layers without scanners
			if layers[s].depth == 0 {
				continue
			}
			// move scanner
			if layers[s].direction == 0 {
				layers[s].scanner++
				if layers[s].scanner == (layers[s].depth - 1) {
					layers[s].direction = 1
				}
			} else {
				layers[s].scanner--
				if layers[s].scanner == 0 {
					layers[s].direction = 0
				}
			}
		}
	}
	/*
		for p := 0; p <= lastID; p++ {
			if layers[p].depth == 0 {
				continue
			}
			fmt.Println(p, layers[p].depth, layers[p].scanner, layers[p].direction)
		}
	*/
	return
}

func main() {
	testCase1 := []string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}
	challengeInput := getChallenge()

	fmt.Println(solve(testCase1) == 24)
	fmt.Println(solve(challengeInput))
}

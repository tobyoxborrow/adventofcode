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
What is the fewest number of picoseconds that you need to delay the packet to
pass through the firewall without being caught?

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

func makeLayers(lines []string) map[int]*layer {
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
	return layers
}

func solve(lines []string) (score int) {
	layers := makeLayers(lines)

	score = 0

	// "packet" movement from left to right
	for p := 0; p < len(layers); p++ {
		// scoring
		if layers[p].scanner == 0 {
			score += layers[p].severity
		}

		// scanner activity
		for s := 0; s < len(layers); s++ {
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
	return
}

type packet struct {
	delay    int
	position int
}

func solveB(lines []string) (delay int) {
	layers := makeLayers(lines)

	packets := make(map[int]*packet)

	// picoseconds
	for ps := 0; ; ps++ {
		// create new packet for this picosecond
		packets[ps] = &packet{ps, 0}

		// packet activity
		for packet := range packets {
			// check collision
			if layers[packets[packet].position].scanner == 0 {
				delete(packets, packet)
				continue
			}
			// movement
			packets[packet].position++

			if packets[packet].position >= len(layers) {
				delay = packets[packet].delay
				return
			}
		}

		// scanner activity
		for s := 0; s < len(layers); s++ {
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

	fmt.Println(solveB(testCase1) == 10)
	fmt.Println(solveB(challengeInput))
}

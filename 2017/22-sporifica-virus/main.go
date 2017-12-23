package main

/*

Day 22: Sporifica Virus

..#
#..
...

A:
How many bursts cause a node to become infected?

B:
How many bursts cause a node to become infected?

*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

// make a 2d grid of bits of zeros
func makeBits(size int) (bits [][]int8) {
	bits = make([][]int8, size)
	for y := 0; y < size; y++ {
		rbits := make([]int8, size)
		bits[y] = rbits
	}
	return
}

type vector struct {
	x int
	y int
}

var (
	up    = vector{0, -1}
	down  = vector{0, +1}
	left  = vector{-1, 0}
	right = vector{+1, 0}
)

const (
	clean    = 0
	infected = 1
	weakened = 2
	flagged  = 3
	modeA    = 0
	modeB    = 1
)

type carrier struct {
	grid          *grid
	position      vector
	direction     vector
	infectedCount int
	mode          int8
}

func newCarrier(grid *grid) *carrier {
	return &carrier{grid, grid.center, up, 0, modeA}
}

func newCarrierB(grid *grid) *carrier {
	return &carrier{grid, grid.center, up, 0, modeB}
}

// turn depending on content of current position
func (s *carrier) turn() {
	switch s.grid.get(s.position) {
	case infected:
		// if the current node is infected
		switch s.direction {
		case up:
			s.direction = right
		case right:
			s.direction = down
		case down:
			s.direction = left
		case left:
			s.direction = up
		}
	case clean:
		// turn right if clean
		switch s.direction {
		case up:
			s.direction = left
		case left:
			s.direction = down
		case down:
			s.direction = right
		case right:
			s.direction = up
		}
	case weakened:
		// nop
	case flagged:
		// reverse direction
		switch s.direction {
		case up:
			s.direction = down
		case left:
			s.direction = right
		case down:
			s.direction = up
		case right:
			s.direction = left
		}
	}
}

func (s *carrier) infect() {
	switch s.grid.get(s.position) {
	case clean:
		// modeA: clean -> infected
		// modeB: clean -> weakened
		if s.mode == modeA {
			s.grid.set(s.position, infected)
			s.infectedCount++
		} else {
			s.grid.set(s.position, weakened)
		}
	case infected:
		// modeA: infected -> clean
		// modeB: infected -> flagged
		if s.mode == modeA {
			s.grid.set(s.position, clean)
		} else {
			s.grid.set(s.position, flagged)
		}
	case weakened:
		// modeB (only): weakened -> infected
		s.grid.set(s.position, infected)
		s.infectedCount++
	case flagged:
		// modeB (only): flagged -> clean
		s.grid.set(s.position, clean)
	}
}

// move the carrier forward
func (s *carrier) move() {
	s.position.x += s.direction.x
	s.position.y += s.direction.y

	// is the grid big enough for where we want to go?
	if !s.isInGrid() {
		// correct our position for the new grid size
		s.position.x += s.grid.size
		s.position.y += s.grid.size
		//fmt.Println(s.position, s.direction)
		s.grid.grow()
	}
}

func (s *carrier) isInGrid() bool {
	if s.position.x < 0 || s.position.x >= s.grid.size || s.position.y < 0 || s.position.y >= s.grid.size {
		return false
	}
	return true
}

type grid struct {
	bits   [][]int8
	size   int
	center vector
}

func (g *grid) print(s string) {
	for y := 0; y < len(g.bits); y++ {
		fmt.Println(s, g.bits[y])
	}
}

// return the value for the node at a given vector on the grid
func (g *grid) get(v vector) int8 {
	return g.bits[v.y][v.x]
}

// set the value for the node at a given vector on the grid
func (g *grid) set(v vector, value int8) {
	g.bits[v.y][v.x] = value
}

// make a larger grid and place the old grid in the center
func (g *grid) grow() {
	//fmt.Println("Growing...")
	size := g.size * 3
	bits := makeBits(size)
	center := vector{size / 2, size / 2}

	// write the old grid onto the new grid
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			bits[y+g.size][x+g.size] = g.bits[y][x]
		}
	}

	g.bits = bits
	g.size = size
	g.center = center
}

// read a pattern from a slice of strings and return a grid representation
func newGrid(lines []string) *grid {
	size := len(lines) // grids are square
	bits := makeBits(size)
	// dump the characters into a 2d grid
	for y, line := range lines {
		row := strings.Split(line, "")
		for x := 0; x < len(row); x++ {
			if row[x] == "#" {
				bits[y][x] = 1
			}
		}
	}
	center := vector{size / 2, size / 2}
	return &grid{bits, size, center}
}

func solve(lines []string, iterations int) int {
	grid := newGrid(lines)

	carrier := newCarrier(grid)

	for c := 0; c < iterations; c++ {
		carrier.turn()
		carrier.infect()
		carrier.move()
	}
	return carrier.infectedCount
}

func solveB(lines []string, iterations int) int {
	grid := newGrid(lines)

	carrier := newCarrierB(grid)

	for c := 0; c < iterations; c++ {
		carrier.turn()
		carrier.infect()
		carrier.move()
	}
	return carrier.infectedCount
}

func main() {
	fmt.Println("A:", solve(getChallenge(), 10000))
	fmt.Println("B:", solveB(getChallenge(), 10000000))
}

package main

/*

Day 22: Sporifica Virus

..#
#..
...

A:
How many bursts cause a node to become infected?

B:
-

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
)

type carrier struct {
	grid          *grid
	position      vector
	direction     vector
	infectedCount int
}

func newCarrier(grid *grid) *carrier {
	return &carrier{grid, grid.center, up, 0}
}

// move the carrier, infecting or cleaning as it goes
func (s *carrier) move() {
	fmt.Println("Moving...")
	fmt.Println(s.position, s.direction)
	node := s.grid.get(s.position)
	// if the current node is infected, turn right otherwise left
	if node == infected {
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
	} else {
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
	}

	// if current node is clean, it becomes infected, otherwise cleaned
	if node == clean {
		s.grid.set(s.position, infected)
		s.infectedCount++
	} else {
		s.grid.set(s.position, clean)
	}

	// move forward
	s.position.x += s.direction.x
	s.position.y += s.direction.y
	fmt.Println(s.position, s.direction)

	// is the grid big enough for where we want to go?
	if !s.isInGrid() {
		// correct our position for the new grid size
		s.position.x += s.grid.size
		s.position.y += s.grid.size
		fmt.Println(s.position, s.direction)
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
	fmt.Println("Growing...")
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
	fmt.Println(grid.size, grid.center)
	grid.print("makegrid")

	carrier := newCarrier(grid)

	for c := 0; c < iterations; c++ {
		carrier.move()
	}
	grid.print("post")
	return carrier.infectedCount
}

func main() {
	fmt.Println("A:", solve(getChallenge(), 10000))
	//fmt.Println("B:", solve(getChallenge(), 10000))
}

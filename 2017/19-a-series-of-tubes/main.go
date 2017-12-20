package main

/*

Day 19: A Series of Tubes

     |
     |  +--+
     A  |  C
 F---|----E|--+
     |  |  |  D
     +B-+  +--+

A:
What letters will it see (in the order it would see them) if it follows the
path?

B:
How many steps does the packet need to go?

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
	return strings.Split(string(b), "\n")
}

var grid [][]string
var steps int

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func changeDirection(y, x int, dir uint) (ny, nx int, ndir uint) {
	ny = y
	nx = x
	// check neighboughs for the next valid move position
	switch {
	case dir != down && grid[y-1][x] != "-" && grid[y-1][x] != " ":
		// go up
		ny--
		ndir = up
	case dir != up && grid[y+1][x] != "-" && grid[y+1][x] != " ":
		// go down
		ny++
		ndir = down
	case dir != left && grid[y][x+1] != "|" && grid[y][x+1] != " ":
		// go right
		nx++
		ndir = right
	case dir != right && grid[y][x-1] != "|" && grid[y][x-1] != " ":
		// go left
		nx--
		ndir = left
	default:
		panic("lost")
	}
	return
}

func walk(y, x int, dir uint) (code string) {
	// where are we?
	switch grid[y][x] {
	case "+":
		ny, nx, ndir := changeDirection(y, x, dir)
		code += walk(ny, nx, ndir)
		// turn in which direction?
	case " ":
		// probably the end
		return
	default:
		if grid[y][x] != "|" && grid[y][x] != "-" && grid[y][x] != " " {
			code += grid[y][x]
		}
		switch dir {
		case up:
			code += walk(y-1, x, up)
		case right:
			code += walk(y, x+1, right)
		case down:
			code += walk(y+1, x, down)
		case left:
			code += walk(y, x-1, left)
		}
	}

	steps++
	return
}

// find the start of the path in the grid
// it is always on the top row
func findStart() (y, x int) {
	y = 1
	for i, v := range grid[y] {
		if v == "|" {
			x = i
		}
	}
	return
}

func makeGrid(lines []string) {
	// grid already created (e.g. go test)
	if len(grid) > 0 {
		return
	}

	// find the grid dimensions
	// not every line is the same width
	// actually that is due to my editor stripping trailing spaces when saving
	// the input but i'll keep this so it is more reliable
	longest := 0
	for _, line := range lines {
		if len(line) > longest {
			longest = len(line)
		}
	}

	paddingFmt := fmt.Sprintf("%%-%ds", longest)
	// add one line to the beginning and end of the slice
	// add one character to the beginning and end of each line
	// so we can skip bounds checking in the walk
	nlines := make([]string, len(lines)+2)
	for i, line := range lines {
		// right pad out the line to match the longest
		pline := fmt.Sprintf(paddingFmt, line)
		nlines[i+1] = " " + pline + " "
	}
	nlines[0] = " " + fmt.Sprintf(paddingFmt, "") + " "
	nlines[len(nlines)-1] = " " + fmt.Sprintf(paddingFmt, "") + " "

	// dump the characters into a 2d grid
	for _, line := range nlines {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
}

func solve(lines []string) string {
	makeGrid(lines)
	starty, startx := findStart()
	code := walk(starty, startx, down)
	return code
}

func solveB(lines []string) int {
	steps = 0
	_ = solve(lines)
	return steps
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	fmt.Println("B:", solveB(getChallenge()))
}

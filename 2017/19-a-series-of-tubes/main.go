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
	// direction: 0 up, 1 right, 2 down, 3 left
	ny = y
	nx = x
	// where can we go within the grid dimensions and where we've been
	okUp := y > 0 && dir != down
	okRight := x < len(grid[y])-1 && dir != left
	okDown := y < len(grid)-1 && dir != up
	okLeft := x > 0 && dir != right
	// check available neighboughs are valid move positions
	switch {
	case okUp && grid[y-1][x] != "-" && grid[y-1][x] != " ":
		// go up
		ny--
		ndir = 0
	case okDown && grid[y+1][x] != "-" && grid[y+1][x] != " ":
		// go down
		ny++
		ndir = 2
	case okRight && grid[y][x+1] != "|" && grid[y][x+1] != " ":
		// go right
		nx++
		ndir = 1
	case okLeft && grid[y][x-1] != "|" && grid[y][x-1] != " ":
		// go left
		nx--
		ndir = 3
	default:
		panic("lost")
	}
	return
}

func walk(y, x int, dir uint) (code string) {
	if y < 0 || y > len(grid) || x < 0 || x > len(grid[0]) {
		panic("left the grid")
	}

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
		case 0:
			code += walk(y-1, x, up)
		case 1:
			code += walk(y, x+1, right)
		case 2:
			code += walk(y+1, x, down)
		case 3:
			code += walk(y, x-1, left)
		}
	}

	steps++
	return
}

// find the start of the path in the grid
// it is always on the top row
func findStart() (y, x int) {
	for i, v := range grid[0] {
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
	// dump the characters into a 2d grid
	paddingFmt := fmt.Sprintf("%%-%ds", longest)
	for _, line := range lines {
		// right pad out the line to match the longest
		pline := fmt.Sprintf(paddingFmt, line)
		tmp := strings.Split(pline, "")
		grid = append(grid, tmp)
	}
}

func solve(lines []string) string {
	makeGrid(lines)
	starty, startx := findStart()
	code := walk(starty, startx, 2)
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

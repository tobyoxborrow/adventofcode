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
	return strings.Split(string(b), "\n")
}

var grid [][]string

func changeDirection(y, x int, dir uint) (ny, nx int, ndir uint) {
	// direction: 0 up, 1 right, 2 down, 3 left
	ny = y
	nx = x
	// where can we go within the grid dimensions and where we've been
	okUp := y > 0 && dir != 2
	okRight := x < len(grid[y])-1 && dir != 3
	okDown := y < len(grid)-1 && dir != 0
	okLeft := x > 0 && dir != 1
	// check available neighboughs are valid move positions
	switch {
	case okUp && grid[y-1][x] != "-" && grid[y-1][x] != " ":
		fmt.Println("going up!", dir)
		// go up
		ny--
		ndir = 0
	case okDown && grid[y+1][x] != "-" && grid[y+1][x] != " ":
		fmt.Println("going down!", dir)
		// go down
		ny++
		ndir = 2
	case okRight && grid[y][x+1] != "|" && grid[y][x+1] != " ":
		fmt.Println("going right!", dir)
		// go right
		nx++
		ndir = 1
	case okLeft && grid[y][x-1] != "|" && grid[y][x-1] != " ":
		fmt.Println("going left!", dir)
		// go left
		nx--
		ndir = 3
	default:
		panic("lost")
	}
	return
}

// direction: 0 up, 1 right, 2 down, 3 left
func walk(y, x int, dir uint) (code string) {
	fmt.Println("walk:", y, x, grid[y][x], dir, len(grid), len(grid[0]))
	if y < 0 || y > len(grid) || x < 0 || x > len(grid[0]) {
		panic("left the grid")
	}

	// where are we?
	switch grid[y][x] {
	case "+":
		fmt.Println("cross:", y, x, dir)
		ny, nx, ndir := changeDirection(y, x, dir)
		fmt.Println("ncross:", ny, nx, ndir)
		code += walk(ny, nx, ndir)
		// turn in which direction?
	case " ":
		// probably the end
		return
	default:
		fmt.Println("road:", y, x)
		if grid[y][x] != "|" && grid[y][x] != "-" && grid[y][x] != " " {
			code += grid[y][x]
		}
		switch dir {
		case 0:
			code += walk(y-1, x, dir)
			fmt.Println("code0:", code)
		case 1:
			code += walk(y, x+1, dir)
			fmt.Println("code1:", code)
		case 2:
			code += walk(y+1, x, dir)
			fmt.Println("code2:", code)
		case 3:
			code += walk(y, x-1, dir)
			fmt.Println("code3:", code)
		}
	}
	return
}

func solve(lines []string) string {
	// find the grid dimensions
	// not every line is the same width
	longest := 0
	for _, line := range lines {
		if len(line) > longest {
			longest = len(line)
		}
	}
	// dump the characters into a 2d grid
	paddingFmt := fmt.Sprintf("%%-%ds", longest)
	fmt.Println("paddingfmt:", paddingFmt)
	for _, line := range lines {
		// right pad out the line to match the longest
		pline := fmt.Sprintf(paddingFmt, line)
		tmp := strings.Split(pline, "")
		grid = append(grid, tmp)
	}

	// find the start
	posy := 0
	posx := 0
	for x, g := range grid[0] {
		if g == "|" {
			posx = x
		}
	}
	fmt.Println(posy, posx)

	// walk the grid
	code := walk(posy, posx, 2)
	fmt.Println(code)
	return code
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	//fmt.Println("B:", solveB(getChallenge()))
}

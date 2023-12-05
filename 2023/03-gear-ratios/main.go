package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	grid := parseInput(input)
	fmt.Println("One:", SolveOne(grid))
	//fmt.Println("Two:", SolveTwo(grid))
}

func parseInput(s string) Grid {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	grid := Grid{
		cells: make(Cells),
	}

	// build grid from input file

	for y, line := range lines {
		for x, item := range line {
			coordinate := Coordinate{
				x: x,
				y: y,
			}
			cell := grid.NewCell(item)
			grid.cells[coordinate] = cell

			grid.width = max(grid.width, x)
			grid.height = max(grid.height, y)
		}
	}

	// find part numbers

	for y := 0; y <= grid.height; y++ {
		for x := 0; x <= grid.width; x++ {
			cell := grid.cells[Coordinate{x, y}]
			if cell.isPart {
				//cell.partNumbers = grid.FindNumbers(Coordinate{x, y})
				//grid.cells[Coordinate{x, y}] = cell
				//fmt.Printf("%s (%d, %d) -> ", string(cell.item), x, y)
				for _, partNumber := range grid.FindNumbers(Coordinate{x, y}) {
					grid.partOne += partNumber
					//fmt.Printf("%d ", partNumber)
				}
				//fmt.Printf("\n")
			}
		}
	}

	return grid
}

type Coordinate struct {
	x int
	y int
}

type Cell struct {
	item     rune // original value of the cell
	isNumber bool
	isPart   bool
	//partNumbers []int
}

type Cells map[Coordinate]Cell

type Grid struct {
	cells   Cells
	width   int
	height  int
	partOne int // answer to Part One
}

func (g *Grid) NewCell(item rune) (cell Cell) {
	cell.item = item
	if item == '.' {
		cell.isNumber = false
		cell.isPart = false
	} else if item >= '0' && item <= '9' {
		cell.isNumber = true
	} else if item >= 33 && item <= 47 {
		cell.isPart = true
	} else if item >= 58 && item <= 64 {
		cell.isPart = true
	} else if item >= 91 && item <= 96 {
		cell.isPart = true
	} else if item >= 123 {
		cell.isPart = true
	}
	return
}

func (g *Grid) FindNumbers(c Coordinate) (numbers []int) {
	inNumber := false
	// upper row
	if g.cells[Coordinate{c.x - 1, c.y - 1}].isNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x - 1, c.y - 1}))
		inNumber = true
	}
	if g.cells[Coordinate{c.x, c.y - 1}].isNumber {
		if !inNumber {
			numbers = append(numbers, g.FindNumber(Coordinate{c.x, c.y - 1}))
		}
		inNumber = true
	} else {
		inNumber = false
	}
	if g.cells[Coordinate{c.x + 1, c.y - 1}].isNumber && !inNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x + 1, c.y - 1}))
	}
	// same row (skipping the part's position)
	if g.cells[Coordinate{c.x - 1, c.y}].isNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x - 1, c.y}))
	}
	if g.cells[Coordinate{c.x + 1, c.y}].isNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x + 1, c.y}))
	}
	// lower row
	inNumber = false
	if g.cells[Coordinate{c.x - 1, c.y + 1}].isNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x - 1, c.y + 1}))
		inNumber = true
	}
	if g.cells[Coordinate{c.x, c.y + 1}].isNumber {
		if !inNumber {
			numbers = append(numbers, g.FindNumber(Coordinate{c.x, c.y + 1}))
		}
		inNumber = true
	} else {
		inNumber = false
	}
	if g.cells[Coordinate{c.x + 1, c.y + 1}].isNumber && !inNumber {
		numbers = append(numbers, g.FindNumber(Coordinate{c.x + 1, c.y + 1}))
	}
	return
}

func (g *Grid) FindNumber(c Coordinate) (number int) {
	// reverse to the start of the number, then read forwards
	start := c
	for g.cells[start].isNumber {
		if start.x < 0 {
			break
		}
		start.x -= 1
	}
	start.x += 1
	tmp := make([]rune, 0)
	for i := start.x; i <= g.width; i++ {
		cell := g.cells[Coordinate{i, c.y}]
		if !cell.isNumber {
			break
		}
		tmp = append(tmp, cell.item)
	}
	number, err := strconv.Atoi(string(tmp))
	if err != nil {
		panic(err)
	}
	return
}

func SolveOne(grid Grid) int {
	return grid.partOne
}

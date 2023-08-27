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

func parseInput(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

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

	lines := parseInput(input)
	locations := newLocations(lines)
	grid := newGrid(locations)
	grid.populateDistances()

	// drawGridA(grid)
	// drawGridB(grid)

	fmt.Println("A:", SolveA(grid))
	fmt.Println("B:", SolveB(grid))
}

type Point struct {
	x int
	y int
}

type Location struct {
	id         int
	point      Point
	isInfinite bool
}

type Locations map[int]Location

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (location *Location) DistanceFrom(point Point) int {
	return abs(location.point.x-point.x) + abs(location.point.y-point.y)
}

type GridItem struct {
	closest     int
	isLocation  bool
	distanceSum int
}

type Grid struct {
	points    []GridItem
	locations Locations
	maxX      int
	maxY      int
}

func newLocations(lines []string) Locations {
	locations := make(Locations)
	for index, line := range lines {
		line = strings.TrimSpace(line)
		splitPos := strings.Index(line, ",")
		x, err := strconv.Atoi(line[0:splitPos])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(line[splitPos+2:])
		if err != nil {
			panic(err)
		}
		locations[index] = Location{
			id:         index,
			point:      Point{x, y},
			isInfinite: false,
		}
	}
	return locations
}

func newGrid(locations Locations) Grid {
	maxX := 0
	maxY := 0
	for _, location := range locations {
		if location.point.x > maxX {
			maxX = location.point.x
		}
		if location.point.y > maxY {
			maxY = location.point.y
		}
	}

	grid := Grid{
		points:    make([]GridItem, 0, maxX*maxY),
		locations: locations,
		maxX:      maxX,
		maxY:      maxY,
	}

	return grid
}

func (grid *Grid) populateDistances() {
	for gy := 0; gy <= grid.maxY+1; gy++ {
		for gx := 0; gx <= grid.maxX+1; gx++ {
			gridItem := GridItem{
				closest:     -1,
				isLocation:  false,
				distanceSum: 0,
			}
			gridPoint := Point{x: gx, y: gy}

			shortestDistance := grid.maxX + grid.maxY
			shortestMatches := 0
			shortestMatch := -1

			for _, location := range grid.locations {
				pointDistance := location.DistanceFrom(gridPoint)
				gridItem.distanceSum += pointDistance
				if pointDistance < shortestDistance {
					shortestDistance = pointDistance
					shortestMatches = 1
					shortestMatch = location.id
				} else if pointDistance == shortestDistance {
					shortestMatches++
				}
			}

			if shortestDistance == 0 {
				gridItem.isLocation = true
			}

			if shortestMatches == 1 {
				gridItem.closest = shortestMatch
				if gx < 1 || gy < 1 || gx > grid.maxX || gy > grid.maxY {
					location := grid.locations[gridItem.closest]
					location.isInfinite = true
					grid.locations[gridItem.closest] = location
				}
			}

			grid.points = append(grid.points, gridItem)
		}
	}
}

/*
func drawGridA(grid Grid) {
	fmt.Printf("Locations:\n%#v\n", grid.locations)
	// fmt.Printf("grid: %v wide x %#v high\n", grid.maxX, grid.maxY)
	for gy := 0; gy <= grid.maxY+1; gy++ {
		for gx := 0; gx <= grid.maxX+1; gx++ {
			item := grid.points[Point{gx, gy}]
			alphaId := item.closest
			if alphaId > 26 {
				alphaId -= 26
			}
			var r rune
			if item.closest == -1 {
				r = '.'
			} else if item.isLocation {
				r = rune(alphaId + 65)
			} else {
				if grid.locations[item.closest].isInfinite {
					r = '_'
				} else {
					r = rune(alphaId + 97)
				}
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func drawGridB(grid Grid) {
	fmt.Printf("Locations:\n%#v\n", grid.locations)
	// fmt.Printf("grid: %v wide x %#v high\n", grid.maxX, grid.maxY)
	for gy := 0; gy <= grid.maxY+1; gy++ {
		for gx := 0; gx <= grid.maxX+1; gx++ {
			item := grid.points[Point{gx, gy}]
			r := '.'
			if item.distanceSum < 10000 {
				r = '#'
			}
			alphaId := item.closest
			if alphaId > 26 {
				alphaId -= 26
			}
			if item.isLocation {
				r = rune(alphaId + 65)
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
*/

func SolveA(grid Grid) int {
	areaSizes := make(map[int]int)
	for _, item := range grid.points {
		if grid.locations[item.closest].isInfinite {
			continue
		}
		areaSizes[item.closest]++
	}

	largestAreaSize := 0
	for _, size := range areaSizes {
		if size > largestAreaSize {
			largestAreaSize = size
		}
	}

	return largestAreaSize
}

func SolveB(grid Grid) int {
	count := 0

	for _, item := range grid.points {
		if item.distanceSum < 10000 {
			count++
		}
	}

	return count
}

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var draw = flag.Bool("draw", false, "draw grids")

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

	if *draw {
		drawGridA(grid)
		drawGridB(grid)
	}

	fmt.Println("A:", SolveA(grid))
	fmt.Println("B:", SolveB(grid, 10000))
}

func parseInput(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
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

type Locations []Location

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
	locations := make(Locations, len(lines))
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
			item := GridItem{
				closest:     -1,
				isLocation:  false,
				distanceSum: 0,
			}
			point := Point{gx, gy}

			shortestDistance := grid.maxX + grid.maxY
			shortestMatches := 0
			shortestMatch := -1

			for index := range grid.locations {
				pointDistance := grid.locations[index].DistanceFrom(point)
				item.distanceSum += pointDistance
				if pointDistance < shortestDistance {
					shortestDistance = pointDistance
					shortestMatches = 1
					shortestMatch = grid.locations[index].id
				} else if pointDistance == shortestDistance {
					shortestMatches++
				}
			}

			if shortestDistance == 0 {
				item.isLocation = true
			}

			if shortestMatches == 1 {
				item.closest = shortestMatch
				if gx < 1 || gy < 1 || gx > grid.maxX || gy > grid.maxY {
					// location := grid.locations[item.closest]
					// location.isInfinite = true
					grid.locations[item.closest].isInfinite = true
				}
			}

			grid.points = append(grid.points, item)
		}
	}
}

func drawGridA(grid Grid) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{grid.maxX + 1, grid.maxY + 1}})

	pointsIndex := 0
	for gy := 0; gy <= grid.maxY+1; gy++ {
		for gx := 0; gx <= grid.maxX+1; gx++ {
			item := grid.points[pointsIndex]
			pointsIndex++

			colour := color.RGBA{255, 255, 255, 255}

			if item.closest >= 0 {
				locationAsByte := uint8(math.Round(float64(item.closest) / float64(len(grid.locations)) * 100))

				colour = color.RGBA{100, 155 + locationAsByte, 155 + locationAsByte, 255}
				if grid.locations[item.closest].isInfinite {
					colour = color.RGBA{110 + locationAsByte, 110 + locationAsByte, 110 + locationAsByte, 255}
				}
			}

			if item.isLocation {
				colour = color.RGBA{0, 0, 0, 255}
			}

			img.Set(gx, gy, colour)
		}
	}

	f, _ := os.Create("grid_a.png")
	png.Encode(f, img)
	fmt.Println("Part A written to grid_a.png")
}

func drawGridB(grid Grid) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{grid.maxX + 1, grid.maxY + 1}})

	pointsIndex := 0
	for gy := 0; gy <= grid.maxY+1; gy++ {
		for gx := 0; gx <= grid.maxX+1; gx++ {
			item := grid.points[pointsIndex]
			pointsIndex++

			colour := color.RGBA{255, 255, 255, 255}

			if item.distanceSum < 10000 {
				colour = color.RGBA{100, 200, 200, 255}
			}

			if item.isLocation {
				colour = color.RGBA{0, 0, 0, 255}
			}

			img.Set(gx, gy, colour)
		}
	}

	f, _ := os.Create("grid_b.png")
	png.Encode(f, img)
	fmt.Println("Part B written to grid_b.png")
}

func SolveA(grid Grid) int {
	areaSizes := make(map[int]int)
	for _, item := range grid.points {
		if item.closest == -1 || grid.locations[item.closest].isInfinite {
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

func SolveB(grid Grid, limit int) int {
	count := 0

	for _, item := range grid.points {
		if item.distanceSum < limit {
			count++
		}
	}

	return count
}

package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	almanac := parseInput(input)
	//fmt.Println("One:", SolveOne(almanac))
	fmt.Println("Two:", SolveTwo(almanac))
}

func parseInput(s string) Almanac {
	almanac := Almanac{
		seeds:       []int{},
		soilMap:     numberMap{},
		fertMap:     numberMap{},
		waterMap:    numberMap{},
		lightMap:    numberMap{},
		tempMap:     numberMap{},
		humidityMap: numberMap{},
		locationMap: numberMap{},
		partOne:     0,
	}

	seedsEnd := strings.Index(s, "\n")
	seeds := strings.Fields(s[7:seedsEnd])
	for _, seed := range seeds {
		seed, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		almanac.seeds = append(almanac.seeds, seed)
	}

	soilMapStart := strings.Index(s, "seed-to-soil map:")
	fertMapStart := strings.Index(s, "soil-to-fertilizer map:")
	waterMapStart := strings.Index(s, "fertilizer-to-water map:")
	lightMapStart := strings.Index(s, "water-to-light map:")
	tempMapStart := strings.Index(s, "light-to-temperature map:")
	humidityMapStart := strings.Index(s, "temperature-to-humidity map:")
	locationMapStart := strings.Index(s, "humidity-to-location map:")

	almanac.soilMap = NewMap(s[soilMapStart:fertMapStart])
	almanac.fertMap = NewMap(s[fertMapStart:waterMapStart])
	almanac.waterMap = NewMap(s[waterMapStart:lightMapStart])
	almanac.lightMap = NewMap(s[lightMapStart:tempMapStart])
	almanac.tempMap = NewMap(s[tempMapStart:humidityMapStart])
	almanac.humidityMap = NewMap(s[humidityMapStart:locationMapStart])
	almanac.locationMap = NewMap(s[locationMapStart:])

	return almanac
}

type numberMap struct {
	list []numberMapItem
}

func (n *numberMap) add(sStart, dStart, length int) {
	item := numberMapItem{
		sStart: sStart,
		dStart: dStart,
		length: length,
	}
	n.list = append(n.list, item)
}

func (n *numberMap) lookup(source int) int {
	for _, item := range n.list {
		if source >= item.sStart && source <= item.sStart+item.length {
			offset := source - item.sStart
			return item.dStart + offset
		}
	}
	return source // any source not mapped correspond to the same number
}

type numberMapItem struct {
	sStart int
	dStart int
	length int
}

type Almanac struct {
	seeds       []int
	soilMap     numberMap
	fertMap     numberMap
	waterMap    numberMap
	lightMap    numberMap
	tempMap     numberMap
	humidityMap numberMap
	locationMap numberMap
	partOne     int // answer to Part One
}

func (a *Almanac) seedToLocation(seed int) int {
	soil := a.soilMap.lookup(seed)
	fert := a.fertMap.lookup(soil)
	water := a.waterMap.lookup(fert)
	light := a.lightMap.lookup(water)
	temp := a.tempMap.lookup(light)
	humidity := a.humidityMap.lookup(temp)
	location := a.locationMap.lookup(humidity)

	return location
}

func NewMap(s string) numberMap {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	nMap := numberMap{}

	for _, line := range lines {
		if strings.Index(line, ":") > 0 {
			continue
		}
		tokens := strings.Fields(line)
		dStart, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		sStart, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(err)
		}
		nMap.add(sStart, dStart, length)
	}

	return nMap
}

func SolveOne(almanac Almanac) int {
	location := math.MaxInt32
	for _, seed := range almanac.seeds {
		seedLocation := almanac.seedToLocation(seed)
		location = min(location, seedLocation)
	}
	return location
}

func SolveTwo(almanac Almanac) int {
	location := math.MaxInt32
	for i := 0; i < len(almanac.seeds); i += 2 {
		for j := 0; j < almanac.seeds[i+1]; j++ {
			seed := almanac.seeds[i] + j
			seedLocation := almanac.seedToLocation(seed)
			location = min(location, seedLocation)
		}
	}
	return location
}

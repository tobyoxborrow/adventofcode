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

	//fmt.Println("One:", Solve(parseInputOne(input)))
	fmt.Println("Two:", Solve(parseInputTwo(input)))
}

func parseInputOne(s string) Races {
	races := Races{}

	lines := strings.Split(s, "\n")
	times := strings.Fields(lines[0][11:])
	distances := strings.Fields(lines[1][11:])

	for index := range times {
		time, err := strconv.Atoi(times[index])
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(distances[index])
		if err != nil {
			panic(err)
		}
		race := Race{
			time:     time,
			distance: distance,
		}
		races.list = append(races.list, race)
	}

	return races
}

func parseInputTwo(s string) Races {
	races := Races{}

	lines := strings.Split(s, "\n")

	time, err := strconv.Atoi(strings.ReplaceAll(lines[0][11:], " ", ""))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.ReplaceAll(lines[1][11:], " ", ""))
	if err != nil {
		panic(err)
	}
	race := Race{
		time:     time,
		distance: distance,
	}
	races.list = append(races.list, race)
	fmt.Printf("%v\n", races)

	return races
}

type Race struct {
	time     int
	distance int
	ways     int
}

func (r *Race) try(holdTime int) {
	moveTime := r.time - holdTime
	if (holdTime * moveTime) > r.distance {
		r.ways += 1
	}
	return
}

type Races struct {
	list []Race
}

func Solve(races Races) (result int) {
	for _, race := range races.list {
		for i := 1; i < race.time; i++ {
			race.try(i)
		}
		if result == 0 {
			result = race.ways
		} else {
			result *= race.ways
		}
	}
	return
}

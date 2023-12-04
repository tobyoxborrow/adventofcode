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

	games := parseInput(input)
	fmt.Println("A:", SolveA(games))
	//fmt.Println("B:", SolveB(games))
}

func parseInput(s string) []Game {
	// Example line:
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	lines := strings.Split(strings.TrimSpace(s), "\n")

	games := make([]Game, 0, len(lines))

	for _, line := range lines {
		colonPosition := strings.Index(line, ":")

		// Game 1:
		tokens := strings.Fields(line[0:colonPosition])
		gameId, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		game := NewGame(gameId)

		// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		sets := strings.FieldsFunc(line[colonPosition+1:], func(r rune) bool { return r == ';' })

		for _, setString := range sets {
			set := Set{
				red:   0,
				green: 0,
				blue:  0,
			}
			// 3 blue, 4 red
			cubes := strings.FieldsFunc(setString, func(r rune) bool { return r == ',' })
			for _, cubeString := range cubes {
				cubeString = strings.TrimSpace(cubeString)
				tokens := strings.Fields(cubeString)
				cubeCount, err := strconv.Atoi(tokens[0])
				if err != nil {
					panic(err)
				}
				switch tokens[1] {
				case "red":
					set.red = cubeCount
				case "green":
					set.green = cubeCount
				case "blue":
					set.blue = cubeCount
				}
			}
			game.AddSet(set)
		}

		games = append(games, game)
	}

	return games
}

type Game struct {
	id    int
	sets  []Set
	value int // if the game is possible, this will be the ID, else zero
}

type Set struct {
	red   int
	green int
	blue  int
}

func NewGame(id int) Game {
	return Game{
		id:    id,
		value: id,
	}
}

func (g *Game) AddSet(set Set) {
	g.sets = append(g.sets, set)

	// if any cubes in this set are impossible, zero the game's value
	if set.red > 12 {
		g.value = 0
	}
	if set.green > 13 {
		g.value = 0
	}
	if set.blue > 14 {
		g.value = 0
	}
}

func SolveA(games []Game) (total int) {
	for _, game := range games {
		total += game.value
	}

	return total
}

package main

import (
	"container/ring"
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
	fmt.Println("A:", SolveA(games[0]))
	fmt.Println("B:", SolveB(games[0]))
}

type Game struct {
	playerCount int
	lastMarble  int
	highScore   int // for sample.txt
}

func parseInput(s string) []Game {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	games := make([]Game, len(lines))

	for index, line := range lines {
		game := Game{}

		tokens := strings.Fields(strings.TrimSpace(line))

		tmp, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		game.playerCount = tmp

		if len(tokens) > 5 {
			tmp, err := strconv.Atoi(tokens[6])
			if err != nil {
				panic(err)
			}
			game.lastMarble = tmp
		}

		if len(tokens) > 10 {
			tmp, err := strconv.Atoi(tokens[11])
			if err != nil {
				panic(err)
			}
			game.highScore = tmp
		}

		games[index] = game
	}

	return games
}

type Player struct {
	id    int
	score int
}

func SolveA(game Game) int {
	players := make([]Player, game.playerCount)
	for i := range players {
		players[i].id = i
	}

	marbleIndex := 0
	r := ring.New(1)
	r.Value = marbleIndex
	for marbleIndex < game.lastMarble {
		for i := range players {
			marbleIndex++
			if marbleIndex > game.lastMarble {
				break
			}
			if marbleIndex%23 == 0 {
				r = r.Move(-9)
				players[i].score += marbleIndex + int(r.Next().Value.(int))
				r.Unlink(1)
				r = r.Move(2)
				continue
			}
			s := ring.New(1)
			s.Value = marbleIndex
			r = r.Link(s)
		}
	}
	highScore := 0
	for i := range players {
		if players[i].score > highScore {
			highScore = players[i].score
		}
	}
	return highScore
}

func SolveB(game Game) int {
	game.lastMarble *= 100
	return SolveA(game)
}

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

	pile := parseInput(input)
	fmt.Println("One:", SolveOne(pile))
	fmt.Println("Two:", SolveTwo(pile))
}

func parseInput(s string) Pile {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	pile := Pile{
		cards:   make([]Card, 0, len(lines)),
		partOne: 0,
	}

	for _, line := range lines {
		card := NewCard(line)
		//fmt.Printf("%v\n", card)
		pile.cards = append(pile.cards, card)
		pile.partOne += card.score
	}

	return pile
}

type numberSet map[int]struct{}

func (n numberSet) add(number int) {
	n[number] = struct{}{}
}

func (n numberSet) has(number int) bool {
	_, ok := n[number]
	return ok
}

type Card struct {
	id      int
	winners numberSet
	numbers numberSet
	score   int
	matches int
}

type Pile struct {
	cards   []Card
	partOne int // answer to Part One
}

func NewCard(line string) Card {
	// Example line:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

	colonPosition := strings.Index(line, ":")
	pipePosition := strings.Index(line, "|")

	// Card ID
	// Card 1:
	tokens := strings.Fields(line[0:colonPosition])
	cardId, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	card := Card{
		id:      cardId,
		winners: make(map[int]struct{}),
		numbers: make(map[int]struct{}),
		score:   0,
		matches: 0,
	}

	// Winning cards
	// 41 48 83 86 17
	winners := strings.Fields(line[colonPosition+1 : pipePosition])

	for _, winnerAsString := range winners {
		winner, err := strconv.Atoi(winnerAsString)
		if err != nil {
			panic(err)
		}
		card.winners.add(winner)
	}

	// Numbers held
	// 83 86  6 31 17  9 48 53
	numbers := strings.Fields(line[pipePosition+1:])

	for _, numberAsString := range numbers {
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			panic(err)
		}
		card.numbers.add(number)
	}

	// Find matches
	score := 0
	for winner := range card.winners {
		if card.numbers.has(winner) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
			card.matches += 1
		}
	}
	card.score += score

	return card
}

func SolveOne(pile Pile) int {
	return pile.partOne
}

func recurseWinners(pile Pile, cardIndex int) (count int) {
	count = 1

	for i := 1; i <= pile.cards[cardIndex].matches; i++ {
		count += recurseWinners(pile, cardIndex+i)
	}

	return
}

func SolveTwo(pile Pile) (count int) {
	for index := range pile.cards {
		count += recurseWinners(pile, index)
	}
	return
}

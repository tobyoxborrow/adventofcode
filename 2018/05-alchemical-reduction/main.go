package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(s string) []byte {
	s = strings.TrimSpace(s)
	return []byte(s)
}

func main() {
	polymer := parseInput(input)

	fmt.Println("A:", SolveA(polymer))
	fmt.Println("B:", SolveB(polymer))
}

type polymer struct {
	stack *list.List
}

func newPolymer() *polymer {
	return &polymer{stack: list.New()}
}

/*
func (p *polymer) Push(unit byte) {
	p.stack.PushFront(unit)
}

func (p *polymer) Pop() {
	if p.stack.Len() > 0 {
		lastUnit := p.stack.Front()
		p.stack.Remove(lastUnit)
	}
}

func (p *polymer) IsEmpty() bool {
	return p.stack.Len() == 0
}
*/

func (p *polymer) Len() int {
	return p.stack.Len()
}

func (p *polymer) Add(unit byte) {
	if p.stack.Len() == 0 {
		p.stack.PushFront(unit)
		return
	}

	lastUnit := p.stack.Front()
	lastValue, _ := lastUnit.Value.(byte)
	if lastValue^unit == 32 {
		p.stack.Remove(lastUnit)
	} else {
		p.stack.PushFront(unit)
	}
}

func SolveA(inputPolymer []byte) int {
	polymer := newPolymer()
	for _, unit := range inputPolymer {
		polymer.Add(unit)
	}
	return polymer.Len()
}

/*
func SolveA(inputPolymer []byte) int {
		i := 0
		for j := 1; j < len(inputPolymer); j++ {
			if inputPolymer[i]^inputPolymer[j] == 32 {
				if i > 0 {
					i--
				}
				continue
			}
			i++
			inputPolymer[i] = inputPolymer[j]
		}
		if i == 0 {
			return 0
		}
		return i + 1
}
*/

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func SolveB(inputPolymer []byte) int {
	shortestPolymer := len(inputPolymer) + 1
	for i := byte('a'); i <= byte('z'); i++ {
		polymer := newPolymer()
		for _, unit := range inputPolymer {
			if unit == i || unit^32 == i {
				continue
			}
			polymer.Add(unit)
		}
		shortestPolymer = min(shortestPolymer, polymer.Len())
	}
	return shortestPolymer
}

package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed sample.txt
var sample string

func TestSolveA(t *testing.T) {
	games := parseInput(sample)

	for _, game := range games {
		name := fmt.Sprintf("Players:%d Last:%d", game.playerCount, game.lastMarble)
		t.Run(name, func(t *testing.T) {
			result := SolveA(game)
			if result != game.highScore {
				t.Fatalf("expected %v, got %v", game.highScore, result)
			}
		})
	}

}

func TestSolveB(t *testing.T) {
	games := parseInput(sample)

	for _, game := range games {
		name := fmt.Sprintf("Players:%d Last:%d", game.playerCount, game.lastMarble)
		t.Run(name, func(t *testing.T) {
			result := SolveA(game)
			if result != game.highScore {
				t.Fatalf("expected %v, got %v", game.highScore, result)
			}
		})
	}
}

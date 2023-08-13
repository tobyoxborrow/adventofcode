package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(s string) []string {
	s = strings.TrimSpace(string(s))
	return strings.Split(s, "\n")
}

func main() {
	lines := parseInput(input)
	claims := makeClaimsFromLines(lines)
	grid := drawGrid(claims)

	fmt.Println(SolveA(grid))
	fmt.Println(SolveB(grid))
}

type Claim struct {
	id      int
	offsetX int
	offsetY int
	width   int
	height  int
}

func parseClaim(s string) Claim {
	/*
	   Example claim:
	   #1 @ 861,330: 20x10
	*/
	tokens := strings.Fields(s)
	claim := Claim{}
	claim.id, _ = strconv.Atoi(tokens[0][1:])
	offsets := strings.FieldsFunc(tokens[2], func(c rune) bool { return c == ',' })
	claim.offsetX, _ = strconv.Atoi(offsets[0])
	claim.offsetY, _ = strconv.Atoi(offsets[1][:len(offsets[1])-1])
	dimensions := strings.FieldsFunc(tokens[3], func(c rune) bool { return c == 'x' })
	claim.width, _ = strconv.Atoi(dimensions[0])
	claim.height, _ = strconv.Atoi(dimensions[1])
	return claim
}

func makeClaimsFromLines(lines []string) []Claim {
	claims := make([]Claim, len(lines))
	for i, line := range lines {
		claims[i] = parseClaim(line)
	}
	return claims
}

type Coordinates struct {
	x int
	y int
}

type Grid map[Coordinates][]Claim

func drawGrid(claims []Claim) Grid {
	grid := make(Grid)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			var emptyClaim []Claim
			grid[Coordinates{i, j}] = emptyClaim
		}
	}

	for _, claim := range claims {
		for claimX := 0; claimX < claim.width; claimX++ {
			for claimY := 0; claimY < claim.height; claimY++ {
				coords := Coordinates{claim.offsetX + claimX, claim.offsetY + claimY}
				grid[coords] = append(grid[coords], claim)
			}
		}
	}

	return grid
}

func SolveA(grid Grid) int {
	overlapped := 0
	for _, v := range grid {
		if len(v) > 1 {
			overlapped++
		}
	}

	return overlapped
}

func SolveB(grid Grid) int {
	seenClaimIds := make(map[int]bool)
	overlappingClaimsIds := make(map[int]bool)

	for _, claimsAtCoords := range grid {
		for _, claim := range claimsAtCoords {
			seenClaimIds[claim.id] = true
			if len(claimsAtCoords) > 1 {
				overlappingClaimsIds[claim.id] = true
			}
		}
	}

	for claimId := range overlappingClaimsIds {
		delete(seenClaimIds, claimId)
	}

	nonOverlappingClaimId := 0

	for claimId := range seenClaimIds {
		nonOverlappingClaimId = claimId
		break
	}

	return nonOverlappingClaimId
}

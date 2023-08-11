package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(s string) []string {
	s = strings.TrimSpace(string(s))
	return strings.Split(s, "\n")
}

func main() {
	fmt.Println(SolveA(input))
	fmt.Println(SolveB(input))
}

func SolveA(input string) int {
	parsed := parseInput(input)

	boxesWithTwo := 0
	boxesWithThree := 0
	for _, entry := range parsed {
		seen := make(map[rune]int)
		for _, letter := range entry {
			count, _ := seen[letter]
			seen[letter] = count + 1
		}

		has_two := false
		has_three := false
		for _, count := range seen {
			if count == 2 {
				has_two = true
				if has_three {
					break
				}
			} else if count == 3 {
				has_three = true
				if has_two {
					break
				}
			}
		}

		if has_two {
			boxesWithTwo++
		}
		if has_three {
			boxesWithThree++
		}

	}

	return boxesWithTwo * boxesWithThree
}

type differenceIndex struct {
	index int
	null  bool
}

func SolveB(input string) string {
	parsed := parseInput(input)

	for i := range parsed {
		for j := i + 1; j < len(parsed); j++ {
			diff := differenceIndex{0, true}
			for k := range parsed[i] {
				if parsed[i][k] != parsed[j][k] {
					if !diff.null {
						diff.null = true
						break
					}
					diff.index = k
					diff.null = false
				}
			}
			if !diff.null {
				return parsed[i][0:diff.index] + parsed[i][diff.index+1:]
			}
		}
	}

	return ""
}

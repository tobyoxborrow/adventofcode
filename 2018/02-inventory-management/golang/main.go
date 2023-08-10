package aoc2018_02

func SolveA(input []string) int {
	boxesWithTwo := 0
	boxesWithThree := 0
	for _, entry := range input {
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

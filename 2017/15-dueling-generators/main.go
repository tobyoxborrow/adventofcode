package main

import (
	"fmt"
	"strconv"
)

/*

Day 15: Dueling Generators

--Gen. A--  --Gen. B--
   1092455   430625591
1181022009  1233683848
 245556042  1431495498
1744312007   137874439
1352636452   285222916

A judge waits for each of them to generate its next value, compares the lowest
16 bits of both values, and keeps track of the number of times those parts of
the values match.

A:
After 40 million pairs, what is the judge's final count?

*/

// common code for both generators
func commonGenerator(prev int, factor int) int {
	magic := 2147483647
	p := prev * factor
	r := p % magic
	return r
}

func generatorAC(seed int) func() int {
	factor := 16807
	prev := seed

	return func() int {
		prev = commonGenerator(prev, factor)
		return prev
	}
}

func generatorBC(seed int) func() int {
	factor := int(48271)
	prev := seed

	return func() int {
		prev = commonGenerator(prev, factor)
		return prev
	}
}

func judge(a, b int) bool {
	// convert to base two
	atwo := fmt.Sprintf("%032s", strconv.FormatInt(int64(a), 2))
	btwo := fmt.Sprintf("%032s", strconv.FormatInt(int64(b), 2))

	// compare final 16 bits
	if atwo[16:] == btwo[16:] {
		return true
	}
	return false
}

func solve(seedA, seedB int) (count int) {
	genA := generatorAC(seedA)
	genB := generatorBC(seedB)
	for c := 0; c < 40000000; c++ {
		a := genA()
		b := genB()

		if judge(a, b) {
			count++
		}
	}
	return
}

func main() {
	fmt.Println(solve(64, 8921) == 588)
	fmt.Println(solve(116, 299))
}
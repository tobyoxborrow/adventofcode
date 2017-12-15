package main

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

import (
	"fmt"
)

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

func generatorA2C(seed int) func() int {
	factor := 16807
	prev := seed

	return func() int {
		for {
			prev = commonGenerator(prev, factor)
			if prev%4 == 0 {
				break
			}
		}
		return prev
	}
}

func generatorB2C(seed int) func() int {
	factor := int(48271)
	prev := seed

	return func() int {
		for {
			prev = commonGenerator(prev, factor)
			if prev%8 == 0 {
				break
			}
		}
		return prev
	}
}

func judge(a, b int) bool {
	// just the last 16 bits
	aa := a & 65535
	bb := b & 65535

	// compare final 16 bits
	return aa == bb
}

func solve(seedA, seedB int) (count int) {
	genA := generatorAC(seedA)
	genB := generatorBC(seedB)
	for c := 0; c < 40000000; c++ {
		if judge(genA(), genB()) {
			count++
		}
	}
	return
}

func solveB(seedA, seedB int) (count int) {
	genA := generatorA2C(seedA)
	genB := generatorB2C(seedB)
	for c := 0; c < 5000000; c++ {
		if judge(genA(), genB()) {
			count++
		}
	}
	return
}

func main() {
	fmt.Println(solve(65, 8921) == 588)
	fmt.Println(solve(116, 299))
	fmt.Println(solveB(65, 8921) == 309)
	fmt.Println(solveB(116, 299))
}

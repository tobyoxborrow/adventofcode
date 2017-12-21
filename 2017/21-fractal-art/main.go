package main

/*

Day 21: Fractal Art

.#.
..#
###

A: How many pixels stay on after 5 iterations?

B:
How many pixels stay on after 18 iterations?

*/

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func getChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

type pattern struct {
	size int
	bits [][]int8
}

// return a pattern based on an input string like ".#./..#/###"
func newPattern(s string) *pattern {
	rows := strings.Split(s, "/")
	size := len(rows[0])
	bits := make([][]int8, size)
	for i, row := range rows {
		rbits := make([]int8, size)
		for j, char := range row {
			switch char {
			case '.':
				rbits[j] = 0
			case '#':
				rbits[j] = 1
			default:
				panic("unknown char")
			}
		}
		bits[i] = rbits
	}
	return &pattern{size, bits}
}

type rule struct {
	needle pattern
	output pattern
}

// returns a new rule created from an input string like "../.. => #.#/.#./..."
func newRule(s string) *rule {
	patterns := strings.Split(s, " => ")
	needle := newPattern(patterns[0])
	output := newPattern(patterns[1])
	return &rule{*needle, *output}
}

var ruleBook []*rule

// parse the "artist's book of enhancement rules"
func readRuleBook(lines []string) []*rule {
	rules := make([]*rule, len(lines))
	for i, line := range lines {
		rpattern := newRule(line)
		rules[i] = rpattern
	}
	return rules
}

func (p *pattern) match(n pattern) (isMatch bool) {
	if n.size != p.size {
		return false
	}
	for y, row := range p.bits {
		for x, elem := range row {
			if elem != n.bits[y][x] {
				return false
			}
		}
	}
	return true
}

// rotate the pattern once clockwise and return the result
// original pattern is unchanged
func (p *pattern) rotate() *pattern {
	bits := make([][]int8, p.size)
	for y := 0; y < p.size; y++ {
		rbits := make([]int8, p.size)
		bits[y] = rbits
	}

	for y := 0; y < p.size; y++ {
		// read a row left-to-right, write a column top-to-bottom
		for x := 0; x < p.size; x++ {
			bits[y][x] = p.bits[p.size-x-1][y]
		}
	}
	return &pattern{p.size, bits}
}

// flip the pattern once horizontally and return the result
// original pattern is unchanged
func (p *pattern) flip() (rotated *pattern) {
	//fmt.Println("flip", p.size)
	//printBits("flip:", p.bits)
	bits := makeBits(p.size)

	for y := 0; y < p.size; y++ {
		// read a row left-to-right, write a row right-to-left
		for x := 0; x < p.size; x++ {
			bits[y][len(p.bits[y])-1-x] = p.bits[y][x]
		}
	}
	return &pattern{p.size, bits}
}

// match pattern (in all rotations) against needles in rule book
// return related output pattern
func (p *pattern) matchRules() pattern {
	//printBits("match:", p.bits)
	// rather than rotate for every rule, save each rotation first
	var rotations [8]*pattern
	rotations[0] = p
	rotations[1] = rotations[0].rotate()
	rotations[2] = rotations[1].rotate()
	rotations[3] = rotations[2].rotate()
	rotations[4] = rotations[0].flip()
	rotations[5] = rotations[1].flip()
	rotations[6] = rotations[2].flip()
	rotations[7] = rotations[3].flip()
	//printBits("rota1:", rotations[1].bits)
	//printBits("flip1:", rotations[5].bits)

	// find a match
	for _, rule := range ruleBook {
		for _, rotation := range rotations {
			if rotation.match(rule.needle) {
				return rule.output
			}
		}
	}
	return pattern{}
}

// divide a pattern into sub patterns
func (p *pattern) divide() (subPatterns []*pattern) {
	var spc int // sub pattern count (how many sub patterns in pattern)
	var sps int // sub pattern size (size of one sub pattern)
	var spw int // sub pattern width (how many per row)
	if p.size%2 == 0 {
		sps = 2
		if p.size > 2 {
			spw = p.size / 2
		} else {
			spw = 1
		}
		spc = spw * spw
	} else {
		sps = 3
		if p.size > 3 {
			spw = p.size / 3
		} else {
			spw = 1
		}
		spc = spw * spw
	}
	if spc == 1 {
		subPatterns = append(subPatterns, p)
		return
	}
	//fmt.Println("divide:", sps, spc, spw)
	for c := 0; c < spc; c++ {
		cyoffset := (c / spw) * sps
		cxoffset := (c % spw) * sps
		//fmt.Println("offsets:", cyoffset, cxoffset)
		sp := makeBits(sps)
		for y := 0; y < sps; y++ {
			for x := 0; x < sps; x++ {
				//fmt.Println("cyx", c, y, x, cyoffset, cxoffset)
				sp[y][x] = p.bits[y+cyoffset][x+cxoffset]
			}
		}
		subPatterns = append(subPatterns, &pattern{sps, sp})
		//printBits("divide:", sp)
	}
	return
}

// takes multiple patterns and joins them together into one, returns result
func joinPatterns(sp []*pattern) *pattern {
	var spc int // sub pattern count
	var sps int // sub pattern size (size of one sub pattern)
	var spw int // sub pattern width (how many per row)
	sps = sp[0].size
	spc = len(sp)
	if spc%2 == 0 {
		spw = int(math.Sqrt(float64(spc)))
	} else {
		spw = int(math.Sqrt(float64(spc)))
	}
	//fmt.Println("join:", sps, spc, spw)
	if spc == 1 {
		return sp[0]
	}
	size := sps * spw
	bits := makeBits(size)
	for c := 0; c < spc; c++ {
		cyoffset := (c / spw) * sps
		cxoffset := (c % spw) * sps
		//fmt.Println("offsets:", cyoffset, cxoffset)
		for y := 0; y < sps; y++ {
			for x := 0; x < sps; x++ {
				bits[y+cyoffset][x+cxoffset] = sp[c].bits[y][x]
			}
		}
		//printBits(fmt.Sprintf("join%d:", c), bits)
	}
	//printBits("join:", bits)
	return &pattern{size, bits}
}

// make a 2d grid of bits of zeros
func makeBits(size int) (bits [][]int8) {
	bits = make([][]int8, size)
	for y := 0; y < size; y++ {
		rbits := make([]int8, size)
		bits[y] = rbits
	}
	return
}

func printBits(s string, bits [][]int8) {
	for i, v := range bits {
		fmt.Println(s, i, v)
	}
}

// enhance 224 176. enhance. stop. move in. stop. pull out, track right. stop.
func (p *pattern) enhance() {
	subPatterns := p.divide()
	fmt.Println("subpatterns count:", len(subPatterns))

	var enhancedSubPatterns = make([]*pattern, 0)
	for _, sp := range subPatterns {
		esp := sp.matchRules()
		//printBits(fmt.Sprintf("esp%d", i), esp.bits)
		enhancedSubPatterns = append(enhancedSubPatterns, &esp)
	}

	enhancedPattern := joinPatterns(enhancedSubPatterns)
	p.size = enhancedPattern.size
	p.bits = enhancedPattern.bits
}

func solve(lines []string, iterations int) (count int) {
	ruleBook = readRuleBook(lines)

	image := newPattern(".#./..#/###")
	for c := 0; c < iterations; c++ {
		image.enhance()
		//printBits(fmt.Sprintf("enhance%d", c), image.bits)
	}

	//printBits("final:", image.bits)

	count = 0
	for _, row := range image.bits {
		for _, v := range row {
			if v == 1 {
				count++
			}
		}
	}
	//fmt.Println(count)

	return
}

func main() {
	fmt.Println("A:", solve(getChallenge(), 5))
	fmt.Println("B:", solve(getChallenge(), 18))
}

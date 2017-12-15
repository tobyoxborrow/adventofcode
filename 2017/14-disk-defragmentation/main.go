package main

/*

Day 14: Disk Defragmentation


*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() string {
	filename := "./input"
	b, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(b))
}

func knotHashC() func([]byte) []int {
	// populate knot list
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	pos := 0
	skip := 0

	return func(lengths []byte) []int {
		for _, v := range lengths {
			l := int(v)
			end := pos + l
			knot := make([]int, end-pos)
			if end < len(list) {
				copy(knot, list[pos:end])
			} else {
				knot = append(list[pos:len(list)], list[0:end-len(list)]...)
			}

			// reverse slice (sort.Reverse didn't seem to work on the append slice)
			// https://github.com/golang/go/wiki/SliceTricks
			for k := len(knot)/2 - 1; k >= 0; k-- {
				opp := len(knot) - 1 - k
				knot[k], knot[opp] = knot[opp], knot[k]
			}

			// write them back to list
			for p, k := pos, 0; k < len(knot); p, k = p+1, k+1 {
				if p >= len(list) {
					p = 0
				}
				list[p] = knot[k]
			}

			pos += l
			pos += skip
			for pos >= len(list) {
				pos -= len(list)
			}
			skip++
		}
		return list
	}
}

// getHash is really solveB from Day 10
func getHash(s string) (hash string) {
	b := []byte(s)

	// add fixed suffix 17, 31, 73, 47, 23
	b = append(b, 17, 31, 73, 47, 23)

	// apply 64 rounds of the knot hash
	knotHash := knotHashC()
	var sh []int // sparse hash
	for i := 0; i < 64; i++ {
		sh = knotHash(b)
	}

	var dh [16]int // dense hash
	for i := 0; i < 16; i++ {
		bx := 0 // block xor value
		for k := 0; k < 16; k++ {
			bx ^= sh[(i*16)+k]
		}
		dh[i] = bx
	}

	// format as hexadecimal string
	hash = ""
	for i := 0; i < 16; i++ {
		hc := strconv.FormatInt(int64(dh[i]), 16)
		hash += fmt.Sprintf("%02s", hc)
	}

	return
}

func rowFromHash(hash string) (row [128]int8) {
	z := 0
	for c := 0; c < len(hash); c++ {
		// convert hex to base ten
		hex := string(hash[c])
		ten, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			panic(err)
		}
		// convert base ten to base two with zero padding
		two := fmt.Sprintf("%04s", strconv.FormatInt(ten, 2))
		//fmt.Println(two)
		for d := 0; d < 4; d++ {
			switch string(two[d]) {
			case "0":
				row[z] = 0
			case "1":
				row[z] = 1
			}
			z++
		}
	}
	return
}

func solve(keyString string) (usedBlocks int) {
	// create and populate grid
	var grid [128][128]int8
	for g := 0; g < len(grid); g++ {
		hashInput := fmt.Sprintf("%s-%d", keyString, g)
		//fmt.Println(hashInput)
		rowHash := getHash(hashInput)
		row := rowFromHash(rowHash)
		//fmt.Println(row)
		grid[g] = row
	}

	// count used blocks
	usedBlocks = 0
	for g := 0; g < len(grid); g++ {
		//fmt.Println(grid[g])
		for r := 0; r < len(grid[g]); r++ {
			if grid[g][r] == 1 {
				usedBlocks++
			}
		}
	}
	return
}

func main() {
	testCase1 := "flqrgnkx"
	challengeInput := getChallenge()
	fmt.Println(solve(testCase1) == 8108)
	fmt.Println(solve(challengeInput))
}

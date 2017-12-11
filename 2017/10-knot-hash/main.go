package main

/*

Day 10: Knot Hash

3, 4, 1, 5

Begin with a list of numbers from 0 to 255, a current position which begins at
0 (the first element in the list), a skip size (which starts at 0), and a
sequence of lengths (your puzzle input). Then, for each length:

* Reverse the order of that length of elements in the list, starting with the
  element at the current position.
* Move the current position forward by that length plus the skip size.
* Increase the skip size by one.

A:
What is the result of multiplying the first two numbers in the list?

B:
Treating your puzzle input as a string of ASCII characters, what is the Knot
Hash of your puzzle input?

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	return s
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

func solve(lengths string) (result int) {
	// turn string into list of byte array
	s := strings.Split(lengths, ",")
	b := make([]byte, len(s))
	for i := range s {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		b[i] = byte(v)
	}

	knotHash := knotHashC()

	list := knotHash(b)

	return list[0] * list[1]
}

func solveB(s string) (hash string) {
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

func main() {
	testCase1 := "12"
	challengeInput := getChallenge()
	fmt.Println(solve(testCase1) == 110)
	fmt.Println(solve(challengeInput))

	testCaseB1 := ""
	testCaseB2 := "AoC 2017"
	testCaseB3 := "1,2,3"
	testCaseB4 := "1,2,4"
	fmt.Println(solveB(testCaseB1) == "a2582a3a0e66e6e86e3812dcb672a272")
	fmt.Println(solveB(testCaseB2) == "33efeb34ea91902bb2f59c9920caa6cd")
	fmt.Println(solveB(testCaseB3) == "3efbe78a8d82f29979031a4aa0b16a9d")
	fmt.Println(solveB(testCaseB4) == "63960835bcdc130f0b66d7ff4f6a5a8e")
	fmt.Println(solveB(challengeInput))
}

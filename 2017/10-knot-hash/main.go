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
-

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	return strings.Split(s, ",")
}

func solve(listSize int, lengths []string) (result int) {
	// populate knot list
	list := make([]int, listSize)
	for i := range list {
		list[i] = i
	}

	fmt.Println("New solution")

	pos := 0
	skip := 0
	// apply length operations on the knot list
	for i, v := range lengths {
		l, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		end := pos + l
		fmt.Println(i, "pos, v, l, end", pos, v, l, end)
		knot := make([]int, end-pos)
		if end < len(list) {
			//fmt.Println(i, "copy")
			copy(knot, list[pos:end])
		} else {
			//fmt.Println(i, "append")
			knot = append(list[pos:len(list)], list[0:end-len(list)]...)
		}
		fmt.Println(i, "knot", len(knot), cap(knot), knot)

		// https://github.com/golang/go/wiki/SliceTricks
		for k := len(knot)/2 - 1; k >= 0; k-- {
			opp := len(knot) - 1 - k
			knot[k], knot[opp] = knot[opp], knot[k]
		}
		//sort.Sort(sort.Reverse(sort.IntSlice(knot)))
		//fmt.Println(i, "knot", len(knot), cap(knot), knot)

		// write them back to list
		for p, k := pos, 0; k < len(knot); p, k = p+1, k+1 {
			if p >= len(list) {
				p = 0
			}
			list[p] = knot[k]
		}

		fmt.Println(i, "list", list)

		pos += l
		pos += skip
		for pos >= len(list) {
			pos -= len(list)
		}
		skip++
		fmt.Println(i, "pos, skip, list[0], list[1]", pos, skip, list[0], list[1])
	}

	fmt.Println(list[0] * list[1])
	return list[0] * list[1]
}

func main() {
	testCase1 := []string{"3", "4", "1", "5"}
	testCase2 := []string{"12"}
	challengeInput := getChallenge()

	fmt.Println(solve(5, testCase1) == 12)
	fmt.Println(solve(255, testCase2) == 110)
	fmt.Println(solve(256, challengeInput))
}

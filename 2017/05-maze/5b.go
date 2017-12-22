package main

/*

Day 5: A Maze of Twisty Trampolines, All Alike

0
3
0
1
-3

After each jump, if the offset was three or more, instead decrease it by 1.
Otherwise, increase it by 1 as before.

Using this rule with the above example, the process now takes 10 steps, and the
offset values after finding the exit are left as 2 3 2 3 -1.

How many steps does it take to reach the exit?

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() []int {
	filename := "./input"
	b, _ := ioutil.ReadFile(filename)
	s := strings.TrimSpace(string(b))
	var i []int
	for _, v := range strings.Fields(s) {
		n, _ := strconv.Atoi(v)
		i = append(i, n)
	}
	return i
}

func solve(instructions []int) (steps int) {
	ilen := len(instructions)
	fmt.Println("Instructions: ", ilen)
	for ip := 0; ip >= 0 && ip < ilen; {
		steps++
		instruction := instructions[ip]
		if instruction >= 3 {
			instructions[ip]--
		} else {
			instructions[ip]++
		}
		ip += instruction
	}
	return
}

func main() {
	fmt.Println(solve([]int{
		0,
		3,
		0,
		1,
		-3,
	}) == 10)
	fmt.Println(solve(getChallenge()))
}

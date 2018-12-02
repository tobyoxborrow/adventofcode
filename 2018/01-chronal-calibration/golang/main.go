package main

/*

Day 01: Chronal Calibration

A:
The device shows a sequence of changes in frequency (your puzzle input). A
value like +6 means the current frequency increases by 6; a value like -3 means
the current frequency decreases by 3.

B:
You notice that the device repeats the same frequency change list over and
over. To calibrate the device, you need to find the first frequency it reaches
twice.

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
	return strings.Split(s, "\n")
}

func solve(lines []string) int {
	frequency := 0
	for _, line := range lines {
		change, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		frequency += change
	}
	return frequency
}

func solveB(lines []string) int {
	history := make(map[int]int8)
	frequency := 0
	for {
		for _, line := range lines {
			// the value stored isn't important, so just using 1
			history[frequency] = 1

			change, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			frequency += change

			_, ok := history[frequency]
			if ok {
				return frequency
			}
		}
	}
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	fmt.Println("B:", solveB(getChallenge()))
}

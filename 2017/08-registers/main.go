package main

/*

Day 8: I Heard You Like Registers

b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10

A:
What is the largest value in any register after completing the instructions in
your puzzle input?

B:
the highest value held in any register during this process

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

type instruction struct {
	register  string
	operation string
	argument  int
	cregister string
	coperator string
	cargument int
}

func parseLine(line string) (i *instruction) {
	// utc dec -736 if p > -7
	fields := strings.Fields(line)
	reg := fields[0]
	op := fields[1]
	arg, _ := strconv.Atoi(fields[2])
	creg := fields[4]
	cop := fields[5]
	carg, _ := strconv.Atoi(fields[6])
	i = &instruction{reg, op, arg, creg, cop, carg}
	return
}

func solve(lines []string) (largestFinalValue int, largestValue int) {
	registers := make(map[string]int)
	largestValue = 0
	for _, line := range lines {
		i := parseLine(line)

		// create any registers mentioned in the instruction
		if _, ok := registers[i.register]; !ok {
			registers[i.register] = 0
		}
		if _, ok := registers[i.cregister]; !ok {
			registers[i.cregister] = 0
		}

		// check condition
		var c bool
		switch i.coperator {
		case "<":
			c = registers[i.cregister] < i.cargument
		case "<=":
			c = registers[i.cregister] <= i.cargument
		case ">":
			c = registers[i.cregister] > i.cargument
		case ">=":
			c = registers[i.cregister] >= i.cargument
		case "==":
			c = registers[i.cregister] == i.cargument
		case "!=":
			c = registers[i.cregister] != i.cargument
		default:
			panic(fmt.Sprintf("Unknown condition %v", i.coperator))
		}

		if !c {
			continue
		}

		// apply operation
		switch i.operation {
		case "inc":
			registers[i.register] += i.argument
		case "dec":
			registers[i.register] -= i.argument
		default:
			panic(fmt.Sprintf("Unknown operation %v", i.operation))
		}

		if registers[i.register] > largestValue {
			largestValue = registers[i.register]
		}
	}

	// find the largest value and return it
	largestFinalValue = 0
	for _, v := range registers {
		if v > largestFinalValue {
			largestFinalValue = v
		}
	}
	return
}

func main() {
	testCase1 := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	challengeInput := getChallenge()

	s1, s2 := solve(testCase1)
	fmt.Println(s1 == 1)
	fmt.Println(s2 == 10)
	s1, s2 = solve(challengeInput)
	fmt.Println(s1)
	fmt.Println(s2)
}

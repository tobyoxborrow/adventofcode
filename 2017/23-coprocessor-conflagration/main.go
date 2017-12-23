package main

/*

Day 23: Coprocessor Conflagration

A:
How many times is the mul instruction invoked?

B:
After setting register a to 1, if the program were to run to completion, what
value would be left in register h?

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
	opcode    string
	argument1 string
	argument2 string
	value1    int
	value2    int
}

func parseInstruction(i string, registers map[string]int) *instruction {
	fields := strings.Fields(i)
	op := fields[0]
	arg1 := fields[1]
	val1, err := strconv.Atoi(fields[1])
	if err != nil {
		val1 = registers[fields[1]]
	}
	var arg2 string
	var val2 int
	if len(fields) > 2 {
		arg2 = fields[2]
		val2, err = strconv.Atoi(fields[2])
		if err != nil {
			val2 = registers[fields[2]]
		}
	}
	return &instruction{op, arg1, arg2, val1, val2}
}

const (
	debugOff = 0
	debugOn  = 1
)

func solve(instructions []string, debugMode int) int {
	var registers = make(map[string]int)
	registers["a"] = debugMode
	mulCount := 0
	ip := 0 // instruction pointer
	for {
		if ip < 0 || ip >= len(instructions) {
			break
		}
		i := parseInstruction(instructions[ip], registers)
		switch i.opcode {
		case "set":
			registers[i.argument1] = i.value2
		case "sub":
			registers[i.argument1] -= i.value2
		case "mul":
			registers[i.argument1] *= i.value2
			mulCount++
		case "jnz":
			if i.value1 != 0 {
				ip += int(i.value2)
				continue
			}
		}
		ip++
	}
	if debugMode == 1 {
		return registers["h"]
	}
	return mulCount
}

func main() {
	fmt.Println("A:", solve(getChallenge(), debugOff))
	fmt.Println("B:", solve(getChallenge(), debugOn))
}

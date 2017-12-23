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
	"math"
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

func solveB(debugMode int) int {
	h := 0 // counter for numbers between start and end that are not prime

	// translated input instructions
	start := 57               // set b 57
	end := start              // set c b
	if debugMode == debugOn { // jnz a 2 & jnz 1 5 (unconditional jump)
		start *= 100    // mul b 100
		start += 100000 // sub b -100000
		end = start     // set c b
		end += 17000    // sub c -17000
	}

	for x := start; x <= end; x += 17 {
		if !isPrimeSqrt(x) {
			h++
		}
	}
	return h
}

// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	fmt.Println("A:", solve(getChallenge(), debugOff))
	fmt.Println("B:", solveB(debugOn))
}

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
	a := debugMode
	var b int // start
	var c int // end
	//d := 0  // loop counter
	//e := 0  // loop counter
	//f := 0  // flag number is not prime
	var g int // general purpose register
	h := 0    // counter for numbers between start and end that are not prime

	// translated input instructions
	b = 57            // set b 57
	c = b             // set c b
	if a == debugOn { // jnz a 2 & jnz 1 5 (unconditional jump)
		b *= 100    // mul b 100
		b += 100000 // sub b -100000
		c = b       // set c b
		c += 17000  // sub c -17000
	}

	for {
		/*
				f = 1 // set f 1
				d = 2 // set d 2
				for {
					e = 2 // set e 2
					for {
						g = d       // set g d
						g *= e      // mul g e
						g -= b      // sub g b
						if g == 0 { // jnz g 2
							f = 0 // set f 0
						}
						e++         // sub e -1
						g = e       // set g e
						g -= b      // sub g b
						if g == 0 { // jnz g -8
							break
						}
					}
					d++         // sub d -1
					g = d       // set g d
					g -= b      // sub g b
					if g == 0 { // jnz g -13
						break
					}
				}
			if f == 0 { // jnz f 2
				h++ // sub h -1
			}
		*/
		if !isPrimeSqrt(b) {
			h++
		}
		g = b       // set g b
		g -= c      // sub g c
		if g == 0 { // jnz g 2
			break // jnz 1 3 (unconditional jump)
		}
		b += 17 // sub b -17
	} // jnz 1 -23 (unconditional jump)
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

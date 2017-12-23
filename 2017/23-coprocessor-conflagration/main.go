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

func solveB(debugMode int) int {
	a := debugMode
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0

	// translated input instructions
	b = 57            // set b 57
	c = b             // set c b
	if a == debugOn { // jnz a 2 & jnz 1 5 (unconditional jump)
		b *= 100    // mul b 100
		b += 100000 // sub b -100000
		c = b       // set c b
		c += 17000  // sub c -17000
	}
	//fmt.Println("start", a, b, c, d, e, f, g, h)
	for {
		f = 1 // set f 1
		d = 2 // set d 2
	i11:
		e = 2 // set e 2
	i12:
		g = d       // set g d
		g *= e      // mul g e
		g -= b      // sub g b
		if g == 0 { // jnz g 2
			f = 0 // set f 0
		}
		e++         // sub e -1
		g = e       // set g e
		g -= b      // sub g b
		if g != 0 { // jnz g -8
			goto i12
		}
		d++         // sub d -1
		g = d       // set g d
		g -= b      // sub g b
		if g != 0 { // jnz g -13
			goto i11
		}
		if f == 0 { // jnz f 2
			h++ // sub h -1
		}
		g = b       // set g b
		g -= c      // sub g c
		if g == 0 { // jnz g 2
			break // jnz 1 3 (unconditional jump)
		}
		b += 17 // sub b -17
	} // jnz 1 -23 (unconditional jump)
	//fmt.Println("end", a, b, c, d, e, f, g, h)
	return h
}

func printRegisters(s string, registers map[byte]int) {
	tmp := s + ": "
	for k, v := range registers {
		tmp += fmt.Sprintf("%s:%d ", string(k), v)
	}
	fmt.Println(tmp)
}

func main() {
	fmt.Println("A:", solve(getChallenge(), debugOff))
	//fmt.Println("B:", solve(getChallenge(), debugOn))
	//fmt.Println("B:", solveB(getChallenge(), debugOn))
}

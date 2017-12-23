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
	var registers = make(map[byte]int)
	registers['a'] = debugMode
	registers['b'] = 0
	registers['c'] = 0
	registers['d'] = 0
	registers['e'] = 0
	registers['f'] = 0
	registers['g'] = 0
	registers['h'] = 0

	// translated input instructions
	registers['b'] = 57             // set b 57
	registers['c'] = registers['b'] // set c b
	if registers['a'] == 0 {        // jnz a 2
		goto i9 // jnz 1 5 (unconditional jump)
	}
	registers['b'] *= 100           // mul b 100
	registers['b'] -= -100000       // sub b -100000
	registers['c'] = registers['b'] // set c b
	registers['c'] -= -17000        // sub c -17000
i9:
	registers['f'] = 1 // set f 1
	registers['d'] = 2 // set d 2
i11:
	registers['e'] = 2 // set e 2
i12:
	registers['g'] = registers['d']  // set g d
	registers['g'] *= registers['e'] // mul g e
	registers['g'] -= registers['b'] // sub g b
	if registers['g'] == 0 {         // jnz g 2
		registers['f'] = 0 // set f 0
	}
	registers['e'] -= -1             // sub e -1
	registers['g'] = registers['e']  // set g e
	registers['g'] -= registers['b'] // sub g b
	if registers['g'] != 0 {         // jnz g -8
		goto i12
	}
	registers['d'] -= -1             // sub d -1
	registers['g'] = registers['d']  // set g d
	registers['g'] -= registers['b'] // sub g b
	if registers['g'] != 0 {         // jnz g -13
		goto i11
	}
	if registers['f'] == 0 { // jnz f 2
		registers['h'] -= -1 // sub h -1
	}
	registers['g'] = registers['b']  // set g b
	registers['g'] -= registers['c'] // sub g c
	if registers['g'] == 0 {         // jnz g 2
		goto end // jnz 1 3
	}
	registers['b'] -= -17 // sub b -17
	goto i9               // jnz 1 -23 (unconditional jump)
end:
	// 0 1 map[b:57 c:57 f:0 d:57 e:57 g:0 h:1 a:0]
	//printRegisters("end", registers)
	//fmt.Println(debugMode, registers['h'], registers)
	return registers['h']
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

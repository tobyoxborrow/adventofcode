package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*

Day 18: Duet

set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2

A:
What is the value of the recovered frequency (the value of the most recently
played sound) the first time a rcv instruction is executed with a non-zero
value?

B:
-

*/

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

func solve(instructions []string) int {
	var registers = make(map[string]int)
	ip := 0 // instruction pointer
	var sounds []int
	for {
		if ip < 0 || ip >= len(instructions) {
			break
		}
		i := parseInstruction(instructions[ip], registers)
		switch i.opcode {
		case "snd":
			sounds = append(sounds, i.value1)
		case "rcv":
			if i.value1 != 0 {
				return sounds[len(sounds)-1]
			}
		case "set":
			registers[i.argument1] = i.value2
		case "add":
			registers[i.argument1] += i.value2
		case "mul":
			registers[i.argument1] *= i.value2
		case "mod":
			registers[i.argument1] %= i.value2
		case "jgz":
			if i.value1 > 0 {
				ip += int(i.value2)
				continue
			}
		}
		ip++
	}
	return 0
}

func execute(instructions []string, cr chan int, cs chan int, pid int) (count int) {
	var registers = make(map[string]int)
	registers["p"] = pid
	count = 0
	ip := 0 // instruction pointer
	for {
		if ip < 0 || ip >= len(instructions) {
			break
		}
		i := parseInstruction(instructions[ip], registers)
		switch i.opcode {
		case "snd":
			cs <- i.value1
			if pid == 1 {
				count++
				counterB <- count
			}
		case "rcv":
			registers[i.argument1] = <-cr
		case "set":
			registers[i.argument1] = i.value2
		case "add":
			registers[i.argument1] += i.value2
		case "mul":
			registers[i.argument1] *= i.value2
		case "mod":
			registers[i.argument1] %= i.value2
		case "jgz":
			if i.value1 > 0 {
				ip += int(i.value2)
				continue
			}
		}
		ip++
	}
	return
}

var counterB = make(chan int, 100)

func solveB(instructions []string) int {
	c0 := make(chan int, 100)
	c1 := make(chan int, 100)
	go execute(instructions, c0, c1, 0)
	go execute(instructions, c1, c0, 1)
	for c := range counterB {
		fmt.Println(c)
	}
	return 0
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	fmt.Println("B:", solveB(getChallenge()))
}

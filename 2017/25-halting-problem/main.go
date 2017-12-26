package main

/*

Day 25: The Halting Problem

A:
Recreate the Turing machine and save the computer! What is the diagnostic
checksum it produces once it's working again?

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
	return strings.Split(string(b), "\n")
}

const (
	left  = -1
	right = 1
)

type stateMachine struct {
	state  byte // current state
	mi     int  // memory index pointer
	steps  int  // number of steps to perform
	states []*state
	memory []int8
}

type state struct {
	substates [2]*substate
}

type substate struct {
	id        int8
	write     int8
	direction int
	next      byte
}

func parseSubState(s []string) *substate {
	var id int8
	var write int8
	var direction int
	var next byte
	if s[0][26] == '0' {
		id = 0
	} else {
		id = 1
	}
	if s[1][22] == '0' {
		write = 0
	} else {
		write = 1
	}
	if s[2][27] == 'l' {
		direction = left
	} else {
		direction = right
	}
	next = s[3][26] - 65 // A -> 0, B -> 1, and so on
	return &substate{id, write, direction, next}
}

func parseBlock(s []string) *state {
	// In state A:
	var substates [2]*substate
	substates[0] = parseSubState(s[1:5])
	substates[1] = parseSubState(s[5:9])
	return &state{substates}
}

func parseSteps(s string) int {
	// Perform a diagnostic checksum after 12173597 steps.
	tokens := strings.Fields(s)
	steps, err := strconv.Atoi(tokens[5])
	if err != nil {
		panic(err)
	}
	return steps
}

func newStateMachine(s []string) *stateMachine {
	// since the input structure is well formed, will make assumptions about
	// value locations
	currentState := s[0][15] - 65 // A -> 0, B -> 1, and so on
	steps := parseSteps(s[1])

	var block []string
	var states []*state
	for c := 3; c < len(s); c++ {
		// use empty lines as signal for new block
		// to help, added a blank line to the end of the test/input rules so
		// the final block parses with the same condition
		if len(s[c]) == 0 {
			if block != nil {
				b := parseBlock(block)
				states = append(states, b)
			}
			block = nil
			continue
		}
		block = append(block, s[c])
	}
	memory := make([]int8, 10)
	return &stateMachine{currentState, 0, steps, states, memory}
}

// read memory value. increase memory if needed
func (sm *stateMachine) read() int8 {
	if sm.mi == len(sm.memory) {
		sm.memory = append(sm.memory, 0)
		return 0 // empty memory is always zero
	} else if sm.mi == -1 {
		// a bit of fiddling
		// don't actually go into negative space, rather, move the memory strip
		// to the right and change what is zero
		sm.memory = append(sm.memory, 0)
		copy(sm.memory[1:], sm.memory)
		sm.memory[0] = 0
		sm.mi = 0
		return 0
	}
	return sm.memory[sm.mi]
}

// write memory value
func (sm *stateMachine) write(value int8) {
	sm.memory[sm.mi] = value
}

func (sm *stateMachine) execute() {
	for c := 0; c < sm.steps; c++ {
		v := sm.read()
		cs := sm.states[sm.state] // current state
		nv := cs.substates[v].write
		sm.write(nv)
		sm.mi += cs.substates[v].direction
		sm.state = cs.substates[v].next
	}
}

func (sm *stateMachine) checksum() (value int) {
	for c := 0; c < len(sm.memory); c++ {
		value += int(sm.memory[c])
	}
	return
}

func solve(lines []string) int {
	sm := newStateMachine(lines)
	sm.execute()
	return sm.checksum()
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
}

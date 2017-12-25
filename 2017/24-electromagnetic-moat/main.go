package main

/*

Day 24: Electromagnetic Moat

A:
What is the strength of the strongest bridge you can make with the components
you have available?

B:
-

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

type component struct {
	id       string
	ports    [2]int
	strength int
}

func parseLine(s string) *component {
	fields := strings.Split(s, "/")
	var ports [2]int
	var err error
	ports[0], err = strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}
	ports[1], err = strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	strength := ports[0] + ports[1]
	return &component{s, ports, strength}
}

var components []*component

func makeComponents(lines []string) (components []*component) {
	components = make([]*component, len(lines))
	for i, line := range lines {
		components[i] = parseLine(line)
	}
	return
}

type bridge struct {
	//visited  map[string]bool
	visited  []string
	strength int
	length   int
}

var bridges []*bridge

func newBridge() *bridge {
	//v := make(map[string]bool)
	v := make([]string, 0)
	return &bridge{v, 0, 0}
}

func addBridge(pbridge *bridge, c *component) *bridge {
	/*
		_visited := make(map[string]bool, len(pbridge.visited)+1)
		for k, v := range pbridge.visited {
			_visited[k] = v
		}
		_visited[c.id] = true
	*/
	_visited := append(pbridge.visited, c.id)
	_strength := pbridge.strength + c.strength
	_length := pbridge.length + 1

	nbridge := &bridge{_visited, _strength, _length}
	bridges = append(bridges, nbridge)
	return nbridge
}

func seenComponent(id string, bridge *bridge) bool {
	for _, v := range bridge.visited {
		if v == id {
			return true
		}
	}
	return false
}

func buildBridges(c *component, pport int, pbridge *bridge) {
	//if pbridge.visited[c.id] {
	if seenComponent(c.id, pbridge) {
		return
	}

	cport := -1
	switch {
	case c.ports[0] == pport:
		cport = c.ports[1]
	case c.ports[1] == pport:
		cport = c.ports[0]
	default:
		return
	}

	cbridge := addBridge(pbridge, c)

	for _, v := range components {
		if v.ports[0] == cport || v.ports[1] == cport {
			buildBridges(v, cport, cbridge)
		}
	}
}

var answerA int
var answerB int

func solve(lines []string) {
	bridges = nil
	components = makeComponents(lines)

	// build some bridges
	for _, c := range components {
		if c.ports[0] == 0 || c.ports[1] == 0 {
			bridge := newBridge()
			buildBridges(c, 0, bridge)
		}
	}

	// which is the best?
	strongest := -1
	longest := -1
	longestStrength := -1
	for _, c := range bridges {
		if c.strength > strongest {
			strongest = c.strength
		}
		if c.length >= longest && c.strength > longestStrength {
			longest = c.length
			longestStrength = c.strength
		}
	}
	answerA = strongest
	answerB = longestStrength
}

func main() {
	solve(getChallenge())
	fmt.Println("A:", answerA)
	fmt.Println("B:", answerB)
}

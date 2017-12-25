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
	path       string
	components []*component
	visited    map[string]bool
	strength   int
}

var bridges []*bridge

func newBridge() *bridge {
	p := ""
	c := make([]*component, 0)
	v := make(map[string]bool)
	return &bridge{p, c, v, 0}
}

func addBridge(pbridge bridge, c *component) bridge {
	_path := pbridge.path
	_components := pbridge.components
	//_visited := pbridge.visited
	_visited := make(map[string]bool, len(pbridge.visited))
	for k, v := range pbridge.visited {
		_visited[k] = v
	}
	_strength := pbridge.strength

	_path += fmt.Sprintf("%s--", c.id)
	_components = append(_components, c)
	_visited[c.id] = true
	_strength += c.strength

	nbridge := &bridge{_path, _components, _visited, _strength}
	bridges = append(bridges, nbridge)
	return *nbridge
}

func buildBridges(c *component, pport int, pbridge bridge) {
	if pbridge.visited[c.id] {
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

func solve(lines []string) int {
	bridges = nil
	components = makeComponents(lines)

	// build some bridges
	for _, c := range components {
		if c.ports[0] == 0 || c.ports[1] == 0 {
			bridge := newBridge()
			buildBridges(c, 0, *bridge)
		}
	}

	// which is the best?
	best := -1
	for _, c := range bridges {
		if c.strength > best {
			best = c.strength
		}
	}
	return best
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	//fmt.Println("B:", solve(getChallenge()))
}

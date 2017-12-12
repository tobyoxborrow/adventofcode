package main

/*

Day 12: Digital Plumber

0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5

A:
How many programs are in the group that contains program ID 0?

B:
How many groups are there in total?

*/

import (
	"fmt"
	"io/ioutil"
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

// basic node lookup table, since the original data is unordered
type basicProgram struct {
	id      string
	program *program
	conns   []string
}

var basicPrograms map[string]basicProgram

type program struct {
	id    string
	conns []*program
}

func newProgram(id string) *program {
	return &program{id, nil}
}

func (p *program) hasConnection(cid string) bool {
	for _, c := range p.conns {
		if c.id == cid {
			return true
		}
	}
	return false
}

func (p *program) addConnections() {
	if p == nil {
		return
	}

	for _, id := range basicPrograms[p.id].conns {
		// ignore connections to self
		if id == p.id {
			continue
		}

		cbp := basicPrograms[id]

		// get or create the program
		conn := cbp.program
		if conn == nil {
			conn = newProgram(id)
			cbp.program = conn
			basicPrograms[id] = cbp
		}

		// skip if a bi-directional connection exists
		if p.hasConnection(id) && conn.hasConnection(p.id) {
			continue
		}

		// add a connection from parent to this child
		if !p.hasConnection(id) {
			p.conns = append(p.conns, conn)
		}

		// add connection from child to parent
		if !conn.hasConnection(p.id) {
			conn.conns = append(conn.conns, p)
		}

		// go deeper
		conn.addConnections()
	}
}

func countPrograms(p *program, visited map[string]bool) (count int) {
	if p == nil {
		return
	}
	visited[p.id] = true
	count++

	for _, c := range p.conns {
		if visited[c.id] {
			continue
		}
		count += countPrograms(c, visited)
	}
	return
}

func parseProgramLine(programLine string) (id string, children []string) {
	fields := strings.Fields(programLine)
	id = fields[0]
	for _, c := range fields[2:] {
		c = strings.TrimSuffix(c, ",")
		children = append(children, c)
	}
	return
}

func solve(lines []string) (count int) {
	// populate basic program lookup table from input lines
	basicPrograms = make(map[string]basicProgram)
	for _, line := range lines {
		id, children := parseProgramLine(line)
		basicPrograms[id] = basicProgram{id, nil, children}
	}

	// build the node tree
	root := newProgram("0")
	root.addConnections()

	// count how many unique nodes in the tree
	visited := make(map[string]bool)
	count = countPrograms(root, visited)
	return
}

func solveB(lines []string) (count int) {
	// populate basic program lookup table from input lines
	basicPrograms = make(map[string]basicProgram)
	for _, line := range lines {
		id, children := parseProgramLine(line)
		basicPrograms[id] = basicProgram{id, nil, children}
	}

	visited := make(map[string]bool)
	for _, v := range basicPrograms {
		// skip programs we've seen before
		_, ok := visited[v.id]
		if ok {
			continue
		}

		// build new node tree
		root := newProgram(v.id)
		root.addConnections()

		// mark all those visited
		_ = countPrograms(root, visited)

		count++
	}
	return
}

func main() {
	testCase1 := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	challengeInput := getChallenge()

	fmt.Println(solve(testCase1) == 6)
	fmt.Println(solve(challengeInput))

	fmt.Println(solveB(testCase1) == 2)
	fmt.Println(solveB(challengeInput))
}

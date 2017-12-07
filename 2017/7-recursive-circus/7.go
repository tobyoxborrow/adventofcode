package main

/*

Day 7: Recursive Circus

pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)

In this example, tknk is at the bottom of the tower (the bottom program), and
is holding up ugml, padx, and fwft. Those programs are, in turn, holding up
other programs.

What is the name of the bottom program?

*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GetChallenge Get the challenge input from the file
func GetChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	return strings.Split(s, "\n")
}

// Solve Return the solution to the input
func Solve(lines []string) (root string) {
	nodes := make(map[string]int)
	for _, line := range lines {
		fields := strings.Fields(line)
		nodeName := fields[0]
		nodes[nodeName]++
		arrowIdx := strings.Index(line, "->")
		if arrowIdx < 0 {
			continue
		}
		subNodes := strings.Split(line[arrowIdx+3:], ", ")
		for _, subNode := range subNodes {
			nodes[subNode]++
		}
	}
	// find the node which appeared once
	for n, v := range nodes {
		if v == 1 {
			root = n
			break
		}
	}
	return
}

func main() {
	fmt.Println(Solve([]string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}) == "tknk")
	fmt.Println(Solve(GetChallenge()))
}

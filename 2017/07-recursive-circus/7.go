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

A:
In this example, tknk is at the bottom of the tower (the bottom program), and
is holding up ugml, padx, and fwft. Those programs are, in turn, holding up
other programs.

What is the name of the bottom program?

B:

Apparently, one program has the wrong weight, and until it's fixed, they're
stuck here.

What would its weight need to be to balance the entire tower?

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

func solveA(lines []string) (root string) {
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

func parseNodeLine(line string) (nodeName string, weight int, children []string) {
	// Example: emlzcpy (106) -> pwmoihf, sdwnkb
	fields := strings.Fields(line)
	nodeName = fields[0]
	weightS := fields[1]
	weight, err := strconv.Atoi(weightS[1 : len(weightS)-1])
	if err != nil {
		panic(err)
	}
	if len(fields) > 2 {
		// skip the field containing the arrow
		// remove trailing comma too
		for _, v := range fields[3:] {
			children = append(children, strings.TrimSuffix(v, ","))
		}
	}
	return
}

// basic node lookup table, since the original data is unordered
type basicNode struct {
	weight   int
	children []string
}

var basicNodes map[string]basicNode

type node struct {
	name     string
	weight   int
	cweight  int // combined weight of children plus itself
	parent   *node
	children []*node
}

func (n *node) addChildren() {
	if n == nil {
		return
	}

	var children []*node
	cweight := n.weight
	for _, name := range basicNodes[n.name].children {
		child := &node{name, basicNodes[name].weight, 0, n, nil}
		child.addChildren()
		children = append(children, child)
		cweight += child.cweight
	}
	n.children = children
	n.cweight = cweight
}

func findUnbalancedNode(n *node) (un *node) {
	if n == nil {
		return
	}

	// count how many times the sub-tower weights occur
	weights := make(map[int]int)
	for _, child := range n.children {
		weights[child.cweight]++
	}
	// more than once is unbalanced
	if len(weights) == 0 {
		return
	}

	// which weight was bad?
	badw := 0
	for wk, wv := range weights {
		if wv == 1 {
			badw = wk
		}
	}
	//fmt.Println("weights, badw:", weights, badw)

	// which node had the bad weight
	var badNode *node
	for _, child := range n.children {
		if child.cweight == badw {
			badNode = child
			break
		}
	}
	// no bad children, we reached the end
	if badNode == nil {
		//fmt.Println("badnode was nil")
		return
	}

	// go deeper. stop when we find no more unbalanced children and return the
	// current child
	un = findUnbalancedNode(badNode)
	if un == nil {
		un = badNode
		//fmt.Println("unbalancednode was nil")
	}
	return
}

func solveB(lines []string) (fixedWeight int) {
	// populate basic node lookup table from input lines
	basicNodes = make(map[string]basicNode)
	for _, line := range lines {
		nodeName, weight, children := parseNodeLine(line)
		basicNodes[nodeName] = basicNode{weight, children}
	}
	// fmt.Println(basicNodes)

	rootName := solveA(lines)
	root := &node{rootName, basicNodes[rootName].weight, 0, nil, nil}
	root.addChildren()
	//fmt.Println("root:", root)

	// starting from the root, find the unbalanced node
	un := findUnbalancedNode(root)
	if un == nil {
		panic("No unbalanced node")
	}
	//fmt.Println("Unblanaced:", un.name)

	// work out the correct weight
	weights := make(map[int]int)
	for _, child := range un.parent.children {
		//fmt.Println(child.name, child.weight, child.cweight)
		weights[child.cweight]++
	}

	// which weight was bad?
	goodw := 0
	badw := 0
	for wk, wv := range weights {
		if wv > 1 {
			goodw = wk
		} else {
			badw = wk
		}
	}

	diff := int(math.Abs(float64(goodw - badw)))
	//fmt.Println("Difference:", diff)
	if goodw > badw {
		fixedWeight = un.weight + diff
	} else {
		fixedWeight = un.weight - diff
	}
	return
}

func main() {
	testCase1 := []string{
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
	}
	testCase2 := []string{
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
		"ugml (56) -> gyxo, ebii, jptl",
		"gyxo (48) -> cnt1, cnt2, cnt3",
		"cntj (57)",
		"cnt1 (7)",
		"cnt2 (7)",
		"cnt3 (7)",
	}
	challengeInput := getChallenge()

	fmt.Println("Part A:")
	fmt.Println(solveA(testCase1) == "tknk")
	fmt.Println(solveA(challengeInput))

	fmt.Println("Part B:")
	fmt.Println(solveB(testCase1) == 60)
	fmt.Println(solveB(testCase2) == 40)
	fmt.Println(solveB(challengeInput))
}

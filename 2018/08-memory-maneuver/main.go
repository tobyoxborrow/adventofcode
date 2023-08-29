package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	numbers := parseInput(input)
	fmt.Println("A:", SolveA(numbers))
	// fmt.Println("B:", SolveB(lines))
}

func parseInput(s string) []byte {
	tokens := strings.Fields(strings.TrimSpace(s))
	bytes := make([]byte, len(tokens))
	for i := 0; i < len(tokens); i++ {
		tmp, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}
		bytes[i] = byte(tmp)
	}
	return bytes
}

type Node struct {
	children      []*Node
	metadataValue int
}

type Tree struct {
	root *Node
}

func NewTree(numbers []byte) Tree {
	rootNode, _ := recurseNumbers(numbers)
	return Tree{
		root: rootNode,
	}
}

func recurseNumbers(numbers []byte) (*Node, int) {
	// A header is always exactly two numbers:
	// - The quantity of child nodes
	// - The quantity of metadata entries
	childrenCount := int(numbers[0])
	metadataCount := int(numbers[1])

	node := Node{}

	numbersIndex := 2

	node.children = make([]*Node, childrenCount)
	for i := 0; i < childrenCount; i++ {
		children, length := recurseNumbers(numbers[numbersIndex:])
		node.children[i] = children
		numbersIndex += length
	}

	for i := 0; i < metadataCount; i++ {
		value := int(numbers[numbersIndex+i])
		node.metadataValue += value
	}
	numbersIndex += metadataCount

	return &node, numbersIndex
}

func (n *Node) SumMetadata() int {
	metadataSum := n.metadataValue

	for i := 0; i < len(n.children); i++ {
		metadataSum += n.children[i].SumMetadata()
	}

	return metadataSum
}

func SolveA(numbers []byte) int {
	tree := NewTree(numbers)

	value := tree.root.SumMetadata()

	return value
}

func SolveB(numbers []byte) int {
	return 0
}

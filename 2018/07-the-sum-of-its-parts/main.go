package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
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

	lines := parseInput(input)

	listA := newAdjacencyList(lines)
	fmt.Println("A:", SolveA(listA))

	listB := newAdjacencyList(lines)
	fmt.Println("B:", SolveB(listB, 5, 60))
}

func parseInput(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

type AdjacencyItem struct {
	outgoing []rune
	incoming []rune
	isActive bool // For Part 2
}

type AdjacencyList map[rune]AdjacencyItem

func newAdjacencyList(lines []string) AdjacencyList {
	list := make(AdjacencyList)
	for _, line := range lines {
		// Step B must be finished before step E can begin.
		keyNode := rune(line[5])
		edgeNode := rune(line[36])
		// keyNode := rune(line[5] - 65)
		// edgeNode := rune(line[36] - 65)

		if entry, ok := list[keyNode]; ok {
			entry.outgoing = append(entry.outgoing, edgeNode)
			list[keyNode] = entry
		} else {
			entry := AdjacencyItem{outgoing: make([]rune, 1, 26), incoming: make([]rune, 0, 26)}
			entry.outgoing[0] = edgeNode
			list[keyNode] = entry
		}
	}

	for keyNode := range list {
		for _, outgoingNode := range list[keyNode].outgoing {
			if entry, ok := list[outgoingNode]; ok {
				entry.incoming = append(entry.incoming, keyNode)
				list[outgoingNode] = entry
			} else {
				entry := AdjacencyItem{outgoing: []rune{}, incoming: make([]rune, 1, 26)}
				entry.incoming[0] = keyNode
				list[outgoingNode] = entry
			}
		}
	}

	return list
}

/*
func (list AdjacencyList) SortedKeys() []rune {
	sortedKeys := make([]rune, 0, len(list))
	for key := range list {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] < sortedKeys[j]
	})
	return sortedKeys
}
*/

func (list AdjacencyList) AvailableKeys() []rune {
	// all keys without incoming neighbors
	// set the initial capacity to the number of keys
	// it probably won't reach that, but won't be any longer
	availableKeys := make([]rune, 0, len(list))
	for key := range list {
		if len(list[key].incoming) == 0 {
			availableKeys = append(availableKeys, key)
		}
	}
	sort.Slice(availableKeys, func(i, j int) bool { return availableKeys[i] > availableKeys[j] })
	return availableKeys
}

/*
func removeNodeFromSlice(s []rune, needle rune) []rune {
	newSlice := make([]rune, 0, len(s)-1)
	for i := 0; i < len(s); i++ {
		if s[i] != needle {
			newSlice = append(newSlice, s[i])
		}
	}
	return newSlice
}
*/

func removeNodeFromSlice(s []rune, needle rune) []rune {
	found := false
	for i, j := 0, 0; i < len(s); i++ {
		if !found && s[i] != needle {
			found = true
			continue
		}
		if found {
			s[j] = s[i]
		}
		j++
	}
	return s[0 : len(s)-1]
}

func (list AdjacencyList) RemoveNode(key rune) {
	// remove the node and any references to it
	incomingNodes := list[key].incoming
	outgoingNodes := list[key].outgoing
	delete(list, key)

	for _, incomingKey := range incomingNodes {
		entry := list[incomingKey]
		entry.outgoing = removeNodeFromSlice(entry.outgoing, key)
		list[incomingKey] = entry
	}

	for _, outgoingKey := range outgoingNodes {
		entry := list[outgoingKey]
		entry.incoming = removeNodeFromSlice(entry.incoming, key)
		list[outgoingKey] = entry
	}
}

func SolveA(list AdjacencyList) string {
	availableKeys := list.AvailableKeys()

	path := make([]rune, 0, len(list))
	for len(availableKeys) > 0 {
		// find the next node to follow
		// first node without any dependancies
		chosenKey := rune(-1)
		for _, key := range availableKeys {
			if len(list[key].incoming) == 0 {
				chosenKey = key
			}
		}
		if chosenKey < 0 {
			break
		}

		path = append(path, chosenKey)

		list.RemoveNode(chosenKey)

		availableKeys = list.AvailableKeys()
	}
	return string(path)
}

type Worker struct {
	node          rune
	timeRemaining int
}

func (w *Worker) HasJob() bool {
	return w.node != 0
}

func (w *Worker) IsFinished() bool {
	return w.node != 0 && w.timeRemaining == 0
}

func (w *Worker) WorkOn(node rune, ticksPerJob int) {
	w.node = node
	w.timeRemaining = int(node) - 65 + ticksPerJob
}

func (w *Worker) Tick() {
	w.timeRemaining--
}

func (w *Worker) Reset() {
	w.node = 0
}

func SolveB(list AdjacencyList, totalWorkers int, ticksPerJob int) int {
	workers := make([]Worker, totalWorkers, totalWorkers)
	for i := 0; i < totalWorkers; i++ {
		workers[i] = Worker{}
	}

	completedJobs := 0
	totalJobs := len(list)

	availableKeys := list.AvailableKeys()

	ticks := -1
	for completedJobs < totalJobs {
		// fmt.Printf("tick: %d completed: %d nodes: %d available: %d workers: %#v\n", ticks, completedJobs, totalJobs, len(availableKeys), workers)
		for workerIndex := range workers {
			// fmt.Printf("worker[%d] contents: %v ", workerIndex, workers[workerIndex].node)
			if workers[workerIndex].HasJob() { // is active
				if !workers[workerIndex].IsFinished() { // still working
					// fmt.Printf(" is working on %s for %d more ticks\n", string(workers[workerIndex].node), workers[workerIndex].timeRemaining)
					workers[workerIndex].Tick()
					continue
				} else { // finished
					// fmt.Printf(" is finished with %s ", string(workers[workerIndex].node))
					list.RemoveNode(workers[workerIndex].node)
					availableKeys = list.AvailableKeys()

					workers[workerIndex].Reset()

					completedJobs++
					if completedJobs == totalJobs {
						// fmt.Printf("(ALL JOBS COMPLETE!)\n")
						break
					}
				}
			}
			// fmt.Printf("is available for work ")

			// the worker is ready for a new job...

			// find the next node to follow
			// first node without any dependancies
			// and not being worked on
			chosenKey := rune(-1)
			for _, key := range availableKeys {
				item := list[key]
				if len(item.incoming) == 0 && !item.isActive {
					chosenKey = key
				}
			}
			if chosenKey < 0 {
				// no work available yet
				// process other workers until work becomes available again
				// fmt.Printf("but no work is available right now!\n")
				continue
			}

			// the worker is ready for a job and there is work available...

			entry := list[chosenKey]
			entry.isActive = true
			list[chosenKey] = entry

			workers[workerIndex].WorkOn(chosenKey, ticksPerJob)
			// fmt.Printf("and will work on %s for %d ticks\n", string(chosenKey), workers[workerIndex].timeRemaining)
		}
		ticks++
	}
	return ticks - 1
}

package main

/*

Day 6: Memory Reallocation

0 2 7 0

The reallocation routine operates in cycles. In each cycle, it finds the memory
bank with the most blocks (ties won by the lowest-numbered memory bank) and
redistributes those blocks among the banks. To do this, it removes all of the
blocks from the selected bank, then moves to the next (by index) memory bank
and inserts one of the blocks. It continues doing this until it runs out of
blocks; if it reaches the last memory bank, it wraps around to the first one.

Given the initial block counts in your puzzle input, how many redistribution
cycles must be completed before a configuration is produced that has been seen
before?

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getChallenge() []int {
	filename := "./input"
	b, _ := ioutil.ReadFile(filename)
	s := strings.TrimSpace(string(b))
	var i []int
	for _, v := range strings.Fields(s) {
		n, _ := strconv.Atoi(v)
		i = append(i, n)
	}
	return i
}

func slotConfiguration(slots []int) (configuration string) {
	for _, v := range slots {
		configuration += strconv.Itoa(v) + ","
	}
	return
}

func solve(slots []int) (cycles int) {
	nslots := len(slots)
	configurations := make(map[string]int8)
	for {
		// compare configuration
		configuration := slotConfiguration(slots)
		//fmt.Println(cycles, configuration)
		_, ok := configurations[configuration]
		if ok {
			break
		}

		// save current configuration
		configurations[configuration] = 1
		cycles++

		// identify largest slot and get its blocks
		largest_i := -1
		largest_v := -1
		for i := 0; i < nslots; i++ {
			if slots[i] > largest_v {
				largest_i = i
				largest_v = slots[i]
			}
		}
		slots[largest_i] = 0

		// redistribute blocks
		i := largest_i
		for v := 0; v < largest_v; v++ {
			if i == (nslots - 1) {
				i = 0
			} else {
				i++
			}
			slots[i]++
		}
	}
	return
}

func main() {
	fmt.Println(solve([]int{
		0,
		2,
		7,
		0,
	}) == 5)
	fmt.Println(solve(getChallenge()))
}

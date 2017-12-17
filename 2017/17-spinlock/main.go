package main

/*

Day 17: Spinlock

A:
What is the value after 2017 in your completed circular buffer?

B:
-

*/

import (
	"fmt"
)

type ring struct {
	data []int
	pos  int
}

func newRing() *ring {
	data := make([]int, 1)
	return &ring{data, 0}
}

func (r *ring) next() {
	r.pos++
	if r.pos >= len(r.data) {
		r.pos = 0
	}
}

func (r *ring) insert(v int) {
	data := make([]int, len(r.data)+1)

	copy(data, r.data[0:r.pos+1])
	data[r.pos+1] = v
	copy(data[r.pos+2:], r.data[r.pos+1:])

	r.data = data
	r.pos++
}

func solve(step int) int {
	buffer := newRing()
	for c := 0; c < 2017; c++ {
		for s := 0; s < step; s++ {
			buffer.next()
		}
		buffer.insert(c + 1)
	}
	return buffer.data[buffer.pos+1]
}

func solveB(step int) int {
	pos := 0
	len := 1
	one := 0
	for c := 0; c < 50e6; c++ {
		pos += step
		for pos >= len {
			pos -= len
		}
		if pos == 0 {
			one = c + 1
		}
		len++
		pos++
	}
	return one
}

func main() {
	fmt.Println(solve(3) == 638)
	fmt.Println(solve(337))
	fmt.Println(solveB(337))
}

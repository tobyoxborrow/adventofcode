package main

/*

Day 11: Hex Ed

ne,ne,ne is 3 steps away.
ne,ne,sw,sw is 0 steps away (back where you started).
ne,ne,s,s is 2 steps away (se,se).
se,sw,se,sw,sw is 3 steps away (s,s,sw).

A:
Determine the fewest number of steps required to reach him.

B:
How many steps away is the furthest he ever got from his starting position?

Based on description here:
https://www.redblobgames.com/grids/hexagons/

*/

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func getChallenge() []string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), ",")
}

type hex struct {
	q float64 // x
	r float64 // y
	s float64 // z
}

func (h *hex) distanceFromZero() uint {
	// hex length
	d := math.Abs(h.q) + math.Abs(h.r) + math.Abs(h.s)
	e := d / 2
	return uint(e)
}

func (h *hex) move(step string) {
	switch step {
	case "n":
		h.r++
		h.s--
	case "nw":
		h.q--
		h.r++
	case "ne":
		h.q++
		h.s--
	case "s":
		h.r--
		h.s++
	case "sw":
		h.q--
		h.s++
	case "se":
		h.q++
		h.r--
	}
}

func solve(steps []string) uint {
	h := &hex{0, 0, 0}
	for _, step := range steps {
		h.move(step)
	}
	return h.distanceFromZero()
}

func solveB(steps []string) uint {
	var furthest uint
	h := &hex{0, 0, 0}
	for _, step := range steps {
		h.move(step)
		d := h.distanceFromZero()
		if d > furthest {
			furthest = d
		}
	}
	return furthest
}

func main() {
	fmt.Println(solve([]string{"ne", "ne", "ne"}) == 3)
	fmt.Println(solve([]string{"ne", "ne", "sw", "sw"}) == 0)
	fmt.Println(solve([]string{"ne", "ne", "s", "s"}) == 2)
	fmt.Println(solve([]string{"se", "sw", "se", "sw", "sw"}) == 3)
	fmt.Println(solve(getChallenge()))
	fmt.Println(solveB(getChallenge()))
}

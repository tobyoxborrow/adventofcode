package main

/*

Day 20: Particle Swarm

For each particle, it provides the X, Y, and Z coordinates for the particle's
position (p), velocity (v), and acceleration (a), each in the format <X,Y,Z>.

p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>

Increase the X velocity by the X acceleration.
Increase the Y velocity by the Y acceleration.
Increase the Z velocity by the Z acceleration.
Increase the X position by the X velocity.
Increase the Y position by the Y velocity.
Increase the Z position by the Z velocity.

A:
Which particle will stay closest to position <0,0,0> in the long term?

B:
-

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
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

type vector struct {
	x float64
	y float64
	z float64
}

type particle struct {
	id int
	p  vector
	v  vector
	a  vector
	d  float64
}

func parseAttribute(attr string) vector {
	attr = strings.TrimSuffix(attr, ">")
	attr = attr[3:len(attr)]
	attr = strings.TrimSpace(attr)
	tokens := strings.Split(attr, ",")
	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(tokens[2])
	if err != nil {
		panic(err)
	}
	return vector{float64(x), float64(y), float64(z)}
}

func solve(lines []string) int {
	var particles []*particle
	for id, line := range lines {
		attrs := strings.Split(line, ", ")
		var p vector
		var v vector
		var a vector
		for _, attr := range attrs {
			switch attr[0] {
			case 'p':
				p = parseAttribute(attr)
			case 'v':
				v = parseAttribute(attr)
			case 'a':
				a = parseAttribute(attr)
			}
		}
		d := math.Abs(p.x) + math.Abs(p.y) + math.Abs(p.z)
		particles = append(particles, &particle{id, p, v, a, d})
	}

	// simulate movement
	// not sure how many rounds, just picked a reasonably big number, turned
	// out to be enough, then scaled it back until it stopped being correct.
	for c := 0; c < 400; c++ {
		for _, particle := range particles {
			particle.v.x += particle.a.x
			particle.v.y += particle.a.y
			particle.v.z += particle.a.z
			particle.p.x += particle.v.x
			particle.p.y += particle.v.y
			particle.p.z += particle.v.z
			particle.d = math.Abs(particle.p.x) + math.Abs(particle.p.y) + math.Abs(particle.p.z)
		}
	}

	// find the closest particle
	closestID := particles[0].id
	closestDistance := particles[0].d
	for _, particle := range particles[1:] {
		if particle.d < closestDistance {
			closestID = particle.id
			closestDistance = particle.d
		}
	}
	return closestID
}

func main() {
	fmt.Println("A:", solve(getChallenge()))
	//fmt.Println("B:", solveB(getChallenge()))
}

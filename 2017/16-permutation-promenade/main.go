package main

/*

Day 16: Permutation Promenade

s1, a spin of size 1: eabcd.
x3/4, swapping the last two programs: eabdc.
pe/b, swapping programs e and b: baedc.

A:
In what order are the programs standing after their dance?

B:
In what order are the programs standing after their billion dances?

*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

func getChallenge() string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

var programs []rune

func spin(count int) {
	// save the elements at the end for later
	end := make([]rune, count)
	copy(end, programs[len(programs)-count:])

	// shift everyone else up
	for i := len(programs) - 1; i >= count; i-- {
		programs[i] = programs[i-count]
	}

	// write the end elements to the front
	for i := 0; i < len(end); i++ {
		programs[i] = end[i]
	}
}

func exchange(a int, b int) {
	programs[a], programs[b] = programs[b], programs[a]
}

func partner(a rune, b rune) {
	// find the indexes for the two elements
	ia := -1
	ib := -1
	for i, v := range programs {
		switch v {
		case a:
			ia = i
		case b:
			ib = i
		}
		if ia >= 0 && ib >= 0 {
			break
		}
	}
	exchange(ia, ib)
}

func dance(sequences []string) {
	for _, s := range sequences {
		if len(s) == 0 {
			continue
		}
		opcode := s[0:1]
		args := strings.Split(s[1:], "/")
		switch opcode {
		case "s":
			arg1, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			spin(arg1)
		case "x":
			arg1, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			arg2, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			exchange(arg1, arg2)
		case "p":
			arg1, _ := utf8.DecodeRuneInString(args[0])
			arg2, _ := utf8.DecodeRuneInString(args[1])
			partner(arg1, arg2)
		}
	}
}

func resetPrograms(newPrograms string) {
	// since we use programs between runs and use it with different sizes, set
	// it to nil and create a new slice of the correct size to avoid previous
	// runs causing problems
	programs = nil
	for _, r := range newPrograms {
		programs = append(programs, r)
	}
	programs = programs[0:len(newPrograms)]
}

func solve(startOrder string, sequenceList string, count int) string {
	resetPrograms(startOrder)
	sequences := strings.Split(sequenceList, ",")
	var prev []string
	for c := 0; c < count; c++ {
		dance(sequences)
		prev = append(prev, string(programs))
		if string(programs) == startOrder {
			r := count % len(prev)
			return prev[r-1]
		}
	}
	return string(programs)
}

func main() {
	fmt.Println(solve("abcde", "s1,x3/4,pe/b,", 1) == "baedc")
	fmt.Println(solve("abcdefghijklmnop", getChallenge(), 1))
	fmt.Println(solve("abcde", "s1,x3/4,pe/b,", 2) == "ceadb")
	fmt.Println(solve("abcdefghijklmnop", getChallenge(), 1e9))
}

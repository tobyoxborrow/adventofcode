package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getChallenge() string {
	filename := "./input"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func solve(s string) (x int) {
	sLen := len(s)
	for c := 0; c < sLen; c++ {
		digit := int(s[c])
		var nextDigit int
		if (c + 1) < sLen {
			nextDigit = int(s[c+1])
		} else {
			nextDigit = int(s[0])
		}
		if digit == nextDigit {
			x += (digit - 48)
		}
	}
	return
}

func main() {
	fmt.Println(solve(getChallenge()))
}

package main

import (
	"fmt"
    "io/ioutil"
    "strings"
)

func GetChallenge() string {
    filename := "./input"
    b, _ := ioutil.ReadFile(filename)
    return strings.TrimSpace(string(b))
}

func Solve(s string) (x int) {
    s_len := len(s)
    for c := 0; c < s_len; c++ {
        digit := int(s[c])
        next_digit := -1
        if (c + 1) < s_len {
            next_digit = int(s[c + 1])
        } else {
            next_digit = int(s[0])
        }
        if digit == next_digit {
            x += (digit - 48)
        }
    }
    return
}

func main() {
    fmt.Println(Solve("1122") == 3)
    fmt.Println(Solve("1111") == 4)
    fmt.Println(Solve(GetChallenge()))
}

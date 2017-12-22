package main

import (
	"fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func GetChallenge() []string {
    filename := "./input"
    b, _ := ioutil.ReadFile(filename)
    s := strings.TrimSpace(string(b))
    return strings.Split(s, "\n")
}

func Solve(s []string) (checksum int64) {
    for c := 0; c < len(s); c++ {
        fields := strings.Fields(s[c])
        lowest, _ := strconv.ParseInt(fields[0], 0, 0)
        highest := lowest
        for f := 0; f < len(fields); f++ {
            digit, _ := strconv.ParseInt(fields[f], 0, 0)
            if digit < lowest {
                lowest = digit
            } else if digit > highest {
                highest = digit
            }
        }
        checksum += (highest - lowest)
    }
    return
}

func main() {
    fmt.Println(Solve([]string{
        "5  1   9   5",
        "7  5   3",
        "2  4   6   8",
    }) == 18)
    fmt.Println(Solve(GetChallenge()))
}

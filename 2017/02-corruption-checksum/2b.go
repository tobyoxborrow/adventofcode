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
        for f := 0; f < len(fields); f++ {
            digit, _ := strconv.ParseInt(fields[f], 0, 0)
            for j := 0; j < len(fields); j++ {
                digit2, _ := strconv.ParseInt(fields[j], 0, 0)
                if digit == digit2 {
                    continue
                }
                result := digit % digit2
                if result == 0 {
                    checksum += digit / digit2
                    break
                }
            }
        }
    }
    return
}

func main() {
    fmt.Println(Solve([]string{
        "5  9   2   8",
        "9  4   7   3",
        "3  8   6   5",
    }) == 9)
    fmt.Println(Solve(GetChallenge()))
}

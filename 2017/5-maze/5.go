package main
/*

Day 5: A Maze of Twisty Trampolines, All Alike

0
3
0
1
-3

(0) 3  0  1  -3  - before we have taken any steps.
(1) 3  0  1  -3  - jump with offset 0 (that is, don't jump at all).
Fortunately, the instruction is then incremented to 1.
2 (3) 0  1  -3  - step forward because of the instruction we just modified. The
first instruction is incremented again, now to 2.
2  4  0  1 (-3) - jump all the way to the end; leave a 4 behind.
2 (4) 0  1  -2  - go back to where we just were; increment -3 to -2.
2  5  0  1  -2  - jump 4 steps forward, escaping the maze.

How many steps does it take to reach the exit?

*/

import (
	"fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func GetChallenge() []int {
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

func Solve(instructions []int) (steps int) {
    ilen := len(instructions)
    fmt.Println("Instructions: ", ilen)
    for ip := 0; ip >= 0 && ip < ilen; {
        steps++
        instruction := instructions[ip]
        instructions[ip]++
        ip += instruction
    }
    return
    /*
    for _, v := range s {
        fields := strings.Fields(v)
        words := make(map[string]int8)
        is_valid := bool(true)
        for _, v := range fields {
            runes := strings.Split(v, "")
            sort.Strings(runes)
            sorted_v := strings.Join(runes, "")
            if _, ok := words[sorted_v]; ok {
                is_valid = bool(false)
                break
            } else {
                words[sorted_v] = 1
            }
        }
        if is_valid {
            valid_count++
        }
    }
    return
    */
}

func main() {
    fmt.Println(Solve([]int{
        0,
        3,
        0,
        1,
        -3,
    }) == 5)
    fmt.Println(Solve(GetChallenge()))
}

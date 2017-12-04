package main
/*

Day 4: High-Entropy Passphrases

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be
rearranged to form the first word.  a ab abc abd abf abj is a valid passphrase,
because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form
any other word.

Under this new system policy, how many passphrases are valid?

*/

import (
	"fmt"
    "io/ioutil"
    "sort"
    "strings"
)

func GetChallenge() []string {
    filename := "./input"
    b, _ := ioutil.ReadFile(filename)
    s := strings.TrimSpace(string(b))
    return strings.Split(s, "\n")
}

func Solve(s []string) (valid_count int) {
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
}

func main() {
    fmt.Println(Solve([]string{
        "abcde fghij",
    }) == 1)
    fmt.Println(Solve([]string{
        "abcde fghij",
        "abcde xyz ecdab",
        "a ab abc abd abf abj",
    }) == 2)
    fmt.Println(Solve(GetChallenge()))
}

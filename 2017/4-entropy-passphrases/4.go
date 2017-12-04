package main
/*

Day 4: High-Entropy Passphrases

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.

How many passphrases are valid?

*/

import (
	"fmt"
    "io/ioutil"
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
            if _, ok := words[v]; ok {
                is_valid = bool(false)
                break
            } else {
                words[v] = 1
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
        "aa bb cc dd ee",
        "aa bb cc dd aa",
        "aa bb cc dd aaa",
    }) == 2)
    fmt.Println(Solve([]string{
        "nyot babgr babgr kqtu kqtu kzshonp ylyk psqk",
        "iix ewj rojvbkk phrij iix zuajnk tadv givslju ewj bda",
        "isjur jppvano vctnpjp ngwzdq pxqfrk mnxxes zqwgnd giqh",
        "ojufqke gpd olzirc jfao cjfh rcivvw pqqpudp",
        "ilgomox extiffg ylbd nqxhk lsi isl nrho yom",
        "feauv scstmie qgbod enpltx jrhlxet qps lejrtxh",
        "wlrxtdo tlwdxor ezg ztp uze xtmw neuga aojrixu zpt",
        "wchrl pzibt nvcae wceb",
        "rdwytj kxuyet bqnzlv nyntjan dyrpsn zhi kbxlj ivo",
        "dab mwiz bapjpz jbzppa",
        "hbcudl tsfvtc zlqgpuk xoxbuh whmo atsxt pzkivuo wsa gjoevr hbcudl",
        "gxhqamx dradmqo gxhqamx gxhqamx",
        "yvwykx uhto ten wkvxyy wdbw",
    }) == 9)
    fmt.Println(Solve(GetChallenge()))
}

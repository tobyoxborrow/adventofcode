#!/usr/bin/env python3
"""

Day 4: High-Entropy Passphrases

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.

How many passphrases are valid?

"""


def get_challenge():
    with open("./input", "r") as input_fh:
        return input_fh.readlines()


def solve(puzzle_input):
    valid_count = 0
    for line in puzzle_input:
        line = line.strip()
        if not line:
            continue
        tokens = line.split(" ")
        words = set()
        is_valid = True
        for token in tokens:
            if token in words:
                is_valid = False
                break
            words.add(token)
        if is_valid:
            valid_count += 1
    return valid_count


print(solve(
    [
        "aa bb cc dd ee",
        "aa bb cc dd aa",
        "aa bb cc dd aaa"
        ]
    ) == 2)
print(solve(
    [
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
        ]
    ) == 9)
print(solve(get_challenge()))

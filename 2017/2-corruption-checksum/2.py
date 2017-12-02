#!/usr/bin/env python3


def get_challenge():
    with open("./input", "r") as input_fh:
        return input_fh.readlines()


def solve(puzzle_input):
    checksum = 0
    for line in puzzle_input:
        line = line.strip()
        if not line:
            continue
        tokens = line.split("\t")
        lowest = None
        highest = None
        for token in tokens:
            digit = int(token)
            if not lowest or digit < lowest:
                lowest = digit
            if not highest or digit > highest:
                highest = digit
        checksum += (highest - lowest)
    return checksum


print(solve(
    [
        "5	1	9	5",
        "7	5	3",
        "2	4	6	8"
        ]
    ) == 18)
print(solve(get_challenge()))
